# O365-Notify

O365 Notify is a simple app to send messages to an Office 365 Group (including a Microsoft Teams channel).

## Webhooks

This utility posts messages to a Office 365 Group or a Teams channel via the use of "Incoming Webhook" Office 365 connector. Check [here](https://docs.microsoft.com/en-us/microsoftteams/platform/concepts/connectors/connectors-using) for instructions on how to implement it on your Teams channel and [here](https://blogs.msdn.microsoft.com/mvpawardprogram/2017/01/17/part-2-office-365-groups/) for an Office 365 Group.
 
## Usage 

You need to setup some env variables in order to use this utility.

- **O365_WEBHOOK**: this is the webhook url (required)
- **O365_MESSAGE**: message you want to send (plain text). This is the message that will be sent if *O365_ADAPTIVECARD* is equal to the empty string
- **O365_ADAPTIVECARD**: the [adaptive card](https://adaptivecards.io) message you want to send. If this env variable is set, the *O365_MESSAGE* is ignored

If you want to run this utility in your shell:

```bash
# will send a simple "Hello world" message
O365_WEBHOOK=https://outlook.office.com/webhook/<GUID>@<GUID>/IncomingWebhook/<GUID>/<GUID> O365_MESSAGE="Hello world" ./o365-notify
```

or

```bash
# will send an adaptive card
O365_WEBHOOK=https://outlook.office.com/webhook/<GUID>@<GUID>/IncomingWebhook/<GUID>/<GUID> O365_ADAPTIVECARD="{\"@type\": \"MessageCard\",\"@context\": \"https:\/\/schema.org\/extensions\",\"summary\": \"Issue 176715375\",\"themeColor\": \"0078D7\",\"title\": \"Issue opened: \\\"Push notifications not working anymore\\\"\",\"sections\": [{\"activityTitle\": \"Miguel Garcie\",\"activitySubtitle\": \"9\/13\/2016, 11:46am\",\"activityImage\": \"https:\/\/connectorsdemo.azurewebsites.net\/images\/MSC12_Oscar_002.jpg\",\"facts\": [{\"name\": \"Repository:\",\"value\": \"mgarcia\\\est\"},{\"name\": \"Issue #:\",\"value\": \"176715375\"}],\"text\": \"There is a problem with Push notifications, they don't seem to be picked up by the connector.\"}],\"potentialAction\": [{\"@type\": \"ActionCard\",\"name\": \"Add a comment\",\"inputs\": [{\"@type\": \"TextInput\",\"id\": \"comment\",\"title\": \"Enter your comment\",\"isMultiline\": true}],\"actions\": [{\"@type\": \"HttpPOST\",\"name\": \"OK\",\"target\": \"http:\/\/...\"}]},{\"@type\": \"HttpPOST\",\"name\": \"Close\",\"target\": \"http:\/\/...\"},{\"@type\": \"OpenUri\",\"name\": \"View in GitHub\",\"targets\": [{\"os\": \"default\",\"uri\": \"http:\/\/...\"}]}]}" ./o365-notify
```

Whereas you should use this command if you want to use the Docker container:

```bash
docker run -e O365_WEBHOOK=https://outlook.office.com/webhook/<GUID>@<GUID>/IncomingWebhook/<GUID>/<GUID> -e O365_MESSAGE="Hello world" dgkanatsios/o365-notify:0.0.1
```

### In Brigade

You can easily use this utility inside Brigade hooks:


```javascript
const {events, Job} = require("brigadier");

events.on("push", (e, p) => {

  var o365 = new Job("o365-notify-notify", "dgkanatsios/o365-notify:0.0.1", ["/o365-notify"]);

  // This doesn't need access to storage, so skip mounting to speed things up.
  o365.storage.enabled = false;
  o365.env = {
    // It's best to store the webhook URL in a project's secrets.
    O365_WEBHOOK: p.secrets.O365_WEBHOOK,
    O365_MESSAGE: "Message Body",
  };
  slack.run();
});
```

---
Inspired by project [technosophos/slack-notify](https://github.com/technosophos/slack-notify).