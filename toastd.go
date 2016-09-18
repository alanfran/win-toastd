/*
  toastd - The Windows toast daemon.

  By Qxcl <qxclpxzk@gmail.com>
  https://github.com/Qxcl/win-toastd
*/

package main

import (
  "github.com/gin-gonic/gin"
  toast "github.com/jacobmarshall/go-toast"
  "log"
  "flag"
  "net"
  "strings"
)

func handler(c *gin.Context) {
  app := c.DefaultQuery("app", "toastd")
  title := c.DefaultQuery("title", "")
  msg := c.DefaultQuery("msg", "")
  icon := c.DefaultQuery("icon", "")

  notification := toast.Notification{
      AppID: app,
      Title: title,
      Message: strings.Replace(msg, "&", "and", -1),
      Icon: icon,
  }

  err := notification.Push()
  if err != nil {
      log.Fatalln(err)
  }
}

func gatekeeper(allowExternal bool) gin.HandlerFunc {
  return func(c *gin.Context) {
    if !allowExternal {
      addresses, err := net.InterfaceAddrs()
      if err != nil {
        c.AbortWithStatus(403)
      }

      addr := c.ClientIP()

      if !contains(addresses, addr) && addr != "::1"{
        c.AbortWithStatus(403)
      }
    }
  }
}

func contains(haystack []net.Addr, needle string) bool {
  c := net.ParseIP(needle)

  for _, ip := range haystack {
    i, _, err := net.ParseCIDR(ip.String())
    if err != nil {panic(err)}

    if i.String() == c.String() {
      return true
    }
  }
  return false
}
func main() {
  port := flag.String("port", "8092", "The port you want to listen to. Default: 8092")
  allowExternal := flag.Bool("allow-external", false, "Allow requests from external IP addresses. {true, false} Default: false")
  flag.Parse()

  gin.SetMode(gin.ReleaseMode)
  r := gin.New()

  r.Use(gatekeeper(*allowExternal))

  r.GET("/", handler)
  r.Run(":"+*port)
}
