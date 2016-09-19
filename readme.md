# toastd - A networked toast daemon for Windows.

toastd lets Windows 10 users receive notifications from another machine. It was designed to work with virtual machine guests, so external IP addresses are ignored by default.

This application listens for requests on port 8092 by default, and it accepts
GET queries and JSON POSTs. Messages containing ampersands (`&`) are not supported at this time, and they will be converted to `+`.

##Parameters:
  * AppID:       The name of the application. Toasts are grouped by app in the notification center.
  * Title:       The title of the toast.
  * Message:     The message body.
  * Icon:        Path to an icon on your Windows system.

###GET Query
Encode the following reserved characters in your query parameters.

| Desired Character |  URL Encoding  | Output |
|:-----------------:|:--------------:|:------:|
| ` ` (space)       | `+` or `%20`   |   ` `  |
| `+`               | `%2b`          |   `+`  |
| `&`               | `%26`          |   `+`  |
| `"`               | `\"`           |   `"`  |
| `!`               | `%21`          |   `!`  |

###POST JSON
You will have to escape `'`, `"`, `\`, and control codes with a backslash `\` (eg: `\'`, `\"`, `\\`). A JSON encoding library should do this for you.

##Example requests:

GET Query:
  ```bash
  curl "192.168.0.5:8092/?AppID=irssi&Title=username&Message=hey+what's+up?&Icon=C:\Users\Username\Icons\irssi.png"
  ```

POST JSON:
  ```bash
  curl -H "Content-Type: application/json" -X POST -d '{"AppID":"irssi", "Title":"username", "Message":"some message", "Icon":"C:/Users/Username/Icons/irssi.png"}' 192.168.0.7:8092
  ```

  ![toast-screenshot](./irssi-notification.png)
##Configuration flags:

  `-port PORTNUMBER`      Changes the listening port.         ex: `win-toastd -port 8082`

  `-allow-external`      Allows requests from external IPs.  ex: `win-toastd -allow-external`

**Warning: Only use -allow-external with a well-configured firewall.**

##Persisting toasts in the Action Center

You must set a registry key for each `AppID` whose notifications you would like to persist in the Action Center.

For example, if you want your irssi notifications to persist, add the key:
`HKEY_CURRENT_USER\SOFTWARE\Microsoft\Windows\CurrentVersion\Notifications\Settings\irssi`
with a DWORD named `ShowInActionCenter` with value `1`.

I will look into automating this with a parameter in the toastd API.

Thanks to "Passing By" for the solution in [this article's](https://deploywindows.info/2015/12/01/powershell-can-i-use-balloons-toasts-and-notifications/) comments, and Mattias Fors for creating the article.
