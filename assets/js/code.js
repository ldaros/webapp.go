function toggleLogGroup(sequenceID) {
  var logGroupLeader = document.getElementById(
    "log-group-leader-" + sequenceID
  );
  var logGroup = document.querySelectorAll(".log-group-" + sequenceID);

  if (logGroupLeader) {
    logGroupLeader.classList.toggle("is-selected");
  }

  logGroup.forEach(function (log) {
    log.classList.toggle("is-hidden");
  });
}
