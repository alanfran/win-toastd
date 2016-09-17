# toastd - A networked toast daemon for Windows.

Designed for Windows users who want to receive notifications from another
machine. Particularly, from within virtual machine guests.

This application listens for requests on a port (default: 8092)
with the following optional query parameters:
  * app:       The name of the application. Toasts are grouped by app in the notification center.
  * title:     The title of the toast.
  * msg:       The message body.
  * icon:      Path to an icon on your Windows system.

You will have to convert your spaces to + or %20 in your scripts. If you
wish to encode a plus sign in your toast, convert it to %2b. This script will
ignore any ampersands in your parameter values.

##Example request:

  ```bash
  curl "192.168.0.5:8092/?app=irssi&title=username&msg=hey+what's+up?&icon=C:\Users\Username\Icons\irssi.png"
  ```

  ![toast-screenshot](./irssi-notification.png)
##Configuration flags:

  `-port PORTNUMBER`      Changes the listening port.         ex: `win-toastd -port 8082`

  `-allow-external`      Allows requests from external IPs.  ex: `win-toastd -allow-external`

**Warning: Only use -allow-external with a well-configured firewall.**
