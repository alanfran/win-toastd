/*
  toastd - The Windows toast daemon.

  By Qxcl <qxclpxzk@gmail.com>
  https://github.com/Qxcl/win-toastd
*/

package main

import (
  "github.com/gin-gonic/gin"
  "github.com/go-toast/toast"
  "flag"
  "net"
)

func handler(c *gin.Context) {
  var notification toast.Notification

  // handles JSON, URL query parameters...
  c.Bind(&notification)

  if notification.AppID == "" {
    notification.AppID = "toastd"
  }

  notification.Push()
}

func gatekeeper(allowExternal bool) gin.HandlerFunc {
  return func(c *gin.Context) {
    if allowExternal { return }

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
  allowExternal := flag.Bool("allow-external", false, "Allow requests from external IP addresses.")
  flag.Parse()

  gin.SetMode(gin.ReleaseMode)
  r := gin.New()

  r.Use(gatekeeper(*allowExternal))

  r.GET("/", handler)
  r.POST("/", handler)
  r.Run(":"+*port)
}
