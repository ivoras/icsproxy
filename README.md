# A HTTP proxy which specifically fixes Outlook ICS (iCal, iCalendar) URLs to have the correct timezone when imported into Google Calendar

This project starts a web server whose only purpose is to fetch a pre-configured ICS URL, probably from Outlook, and serves it modified so that when it's read by Google Calendar, its entries appear in the correct timezone.

The project currently works for the CEST timezone but I intend to make it generic so it can work with arbitrary timezones.

You will, of course, need a way to host this server, either manually or as a Docker container.
