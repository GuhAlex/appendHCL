package main

import (
  "fmt"
  "bufio"
	"os"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)


func main(){

var err error
var qtd int
var terraformFile string
i := 0
fmt.Print("How many paramters to create?:  ")
_, err = fmt.Scanln(&qtd)
if err != nil {
  fmt.Println("error")
}

fmt.Print("PATH of Terraform file to append:  ")
_, err = fmt.Scanln(&terraformFile)
if err != nil {
  fmt.Println("error")
}

for i < qtd {
  nameResource, nameParameter, description, Parametertype, value, env := getParameterValues()
  createResource(terraformFile, nameResource, nameParameter, description, Parametertype, value, env)
  i += 1

  if i >= 1 && i < qtd {
    fmt.Println("Parameters ",i+1,":")
  }

 }
}

func getParameterValues()(nameResource, nameParameter, description, Parametertype, value, env string){

  inputRead := bufio.NewScanner(os.Stdin)

  fmt.Print("Resource Name:  ")
  inputRead.Scan()
  nameResource = inputRead.Text()

  fmt.Print("Parameter Name:  ")
  inputRead.Scan()
  nameParameter = inputRead.Text()

  fmt.Print("Description:  ")
  inputRead.Scan()
  description = inputRead.Text()

  fmt.Print("Type of Parameter[String][SecureString][StringList]:  ")
  inputRead.Scan()
  Parametertype = inputRead.Text()

  fmt.Print("Parameter Value:  ")
  inputRead.Scan()
  value = inputRead.Text()

  fmt.Print("Tags:  ")
  inputRead.Scan()
  env = inputRead.Text()

return nameResource, nameParameter, description, Parametertype, value, env
}

func createResource(terraformFile, nameResource, nameParameter, description, Parametertype, value, env string){
  hclFile := hclwrite.NewEmptyFile()

  filename := terraformFile
  tfFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
    fmt.Println(err)
    return
  }

  rootBody := hclFile.Body()
  // rootBody.AppendNewline()
  rs := rootBody.AppendNewBlock("resource", []string{"aws_ssm_parameter", fmt.Sprintf(nameResource)})
  rsBody := rs.Body()
  rsBody.SetAttributeValue("name", cty.StringVal(fmt.Sprint(nameParameter)))
  rsBody.SetAttributeValue("description", cty.StringVal(fmt.Sprint(description)))
  rsBody.SetAttributeValue("type", cty.StringVal(fmt.Sprint(Parametertype)))
  rsBody.SetAttributeValue("value", cty.StringVal(fmt.Sprint(value)))

  tagBlock := rsBody.AppendNewBlock("tags", nil)
  tagBlockBody := tagBlock.Body()
  tagBlockBody.SetAttributeValue("environment", cty.StringVal(fmt.Sprint(env)))
  rootBody.AppendNewline()

  tfFile.Write(hclFile.Bytes())
}
