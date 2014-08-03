# Pad

Pad is a service that generates anonymous, consumable, expiring codes for use by other services.

For example, say you want to set up an invitation system for your app. You want to generate a code and send it to someone, to verify that they were in fact the person you intended. Pad lets you do that without worrying about generating those codes or worrying about when they expire.

You generate a code, or "pad", and specify how many times it can be consumed, how long it should last before it disappears, and the URL to redirect to when it is consumed. You tell the invitee the URL on Pad to go to in order to consume the pad. If they get there before the pad expires, they are redirected back to your service (via the callback you specified) and you know that they successfully consumed your code.

# Rationale

I'm working on an app that needs this feature. I've worked on several apps that need this feature. I've decided to just make it once, and use it later in whatever place I need it.

There's no reason these invitation codes _must_ be owned by the service that uses them. This prevents coupling the generation and consumption/expiration management of these codes from the rest of your system.

# API

Pad has a simple interface -- you can create pads, and you can consume pads.

## Create Pad

```
POST /pad
```

The body of the POST message should be JSON in the following format:

```
{
  "Consumable": 1, // -1 means there is no limit to the number of times it can be consumed
  "ExpiresInSeconds": 60, // must be positive, maximum 3600 (these are shortlived codes)
  "Callback": "http://yourapp/page?pad={pad}" // the string {pad} will be replaced with your pad
}
```

The server will respond with a JSON object in the same format, but including an `Id` field that contains the value of your pad. You should save this on your end, so you can close the loop on this pad once it is consumed.

## Consume Pad

```
GET /pad/{id}
```

This should be done from the user's browser. Send them to the URL on Pad, and they will consume this code and be sent along back to your service at the callback URL you specified when the pad was created.

# Deployment Concerns

To deploy Pad, you will need to compile it for your production environment. For example, to cross-compile for a 64-bit Linux server:

```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o pad.linux
```

## /etc/pad/pad.gcfg

This file needs to exist to tell Pad how to run.

```
[Host]
port = :80
path = /path/to/pad
```

## systemd config

Edit `/lib/systemd/system/pad.service`:

```
[Unit]
Description=pad web service
After=syslog.target network.target

[Service]
Type=simple
ExecStart=/path/to/pad/pad.linux

[Install]
WantedBy=multi-user.target
```

Then, make a symbolic link:

```
ln -s /lib/systemd/system/pad.service /etc/systemd/system/pad.service
```

And start/enable your service:

```
systemctl start pad.service
systemctl enable pad.service
```
