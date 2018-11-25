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
  o365.run();
});