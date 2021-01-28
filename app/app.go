package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Gerar vai retornar a aplicaçao cli
func Gerar() *cli.App {
   app := cli.NewApp()
   app.Name = "Application CLI"
   app.Usage = "Search IPs and Server names on internet"

   flags :=  []cli.Flag{
	   cli.StringFlag{
	   Name: "host"	,
	   Value: "devbook.com.br",
      },
    }

    app.Commands = []cli.Command{
	{
	   Name: "ip",
	   Usage: "Search IPs and Server names on internet",
	   Flags: flags,
	   Action: SearchIps,
	 },
	 {
		 Name: "servers",
		 Usage: "Search names serve on internet",
		 Flags: flags,
		 Action: SearchServer,
	 },
   }

   return app
}

func SearchIps(c *cli.Context){
   host := c.String("host")

   ips, erro := net.LookupIP(host)

   if erro != nil {
	   log.Fatal(erro)
   }

   for _, ip := range ips {
	   fmt.Println(ip)
   }
}

func SearchServer(c *cli.Context) {
   host := c.String("host")

   servers, erro := net.LookupNS(host)

   if erro != nil {
	   log.Fatal(erro)
   }

   for _, server := range servers {
	   fmt.Println(server)
   }
}