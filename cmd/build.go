package cmd

import (
	"os"
	"io/ioutil"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/juicemia/steve/print"
	"gopkg.in/russross/blackfriday.v2"
)

type Site struct {
	Pages map[string]Page
	Links []string
	CurrentPage Page
}

type Page struct {
	Content string
}

func init() {
	rootCmd.AddCommand(newBuildCmd())
}

func newBuildCmd() *cobra.Command {
	tmpl, err := template.ParseFiles("template/site-content.tpl.html")
	if err != nil {
		print.Fatalf("error reading template file: %v\n", err)
	}

	var srcDir string
	s := &Site{
		Pages: make(map[string]Page),
	}

	var walkFn filepath.WalkFunc
	walkFn = func(path string, info os.FileInfo, err error) error {
		if path == "site" {
			return nil
		}

		print.Verbosef("processing %v...\n", path)

		outpath := fmt.Sprintf("www/%s", path[len(srcDir)+1:])
		if info.Mode().IsDir() {
			if stat, err := os.Stat(outpath); os.IsNotExist(err) {
				print.Verbosef("%v not found, creating directory...", outpath)

				err := os.MkdirAll(outpath, 0644)
				if err != nil {
					return fmt.Errorf("error creating directory %v: %v", outpath, err)
				}
			} else if !stat.Mode().IsDir() {
				return fmt.Errorf("%v already exists and is not a directory", outpath)
			}

			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}

		page := Page{
			Content: string(blackfriday.Run(buf)),
		}

		key := strings.TrimSuffix(strings.TrimPrefix(path, "site/"), ".md") + ".html"
		s.Links = append(s.Links, key)
		s.Pages[key] = page

		fmt.Printf("site in walk fn: %+v\n", s)

		return nil
	}

	cmd := &cobra.Command{
		Use:   "build",
		Short: "Build the blog",
		Long: `Build the blog.

		TODO: add more technical documentation`,
		Run: func(cmd *cobra.Command, args []string) {
			print.Verboseln("building blog...")

			if stat, err := os.Stat("www"); os.IsNotExist(err) {
				err := os.Mkdir("www", 0644)
				if err != nil {
					print.Fatalf("error generating www: %v\n", err)
				}
			} else if !stat.Mode().IsDir() {
				print.Fatalln("www/ exists and is not a directory")
			}

			err := filepath.Walk(srcDir, walkFn)
			if err != nil {
				print.Fatalln(err.Error())
			}

			fmt.Printf("pages: %+v\n", s.Pages)

			for outpath, page := range s.Pages {
                outpath = "www/" + outpath
				print.Verbosef("generating html file at %v...\n", outpath)
				if _, err := os.Stat(outpath); os.IsNotExist(err) {
					f, err := os.Create(outpath)
					if err != nil {
						print.Fatalf("error generating %s: %v", outpath, err)
					}

					s.CurrentPage = page
					err = tmpl.Execute(f, s)
					if err != nil {
						print.Fatalf("error executing template: %v\n", err)
					}
				}
			}
		},
	}

	cmd.PersistentFlags().StringVarP(
		&srcDir,
		"src-dir",
		"s",
		"site",
		"path where the static site content is located",
	)

	return cmd
}


//func generateBasic() {
	//fmt.Println("building site at test-blog/")

	//f, err := os.Open("./test-blog/test.md")
	//if err != nil {
		//panic(err)
	//}

	//buf, err := ioutil.ReadAll(f)
	//if err != nil {
		//panic(err)
	//}

	//output := blackfriday.Run(buf)
	//fmt.Printf("\n\n%s\n", output)
//}

