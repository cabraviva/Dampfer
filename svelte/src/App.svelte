<script lang="ts">
  import "./app.css";
  import LoginCheck from "./popups/LoginCheck.svelte";
  import Sidebar from "./Sidebar.svelte";
  import Page from "./Page.svelte";
  import type { AlertType } from "./types";
  import { dockerReady } from "./script/api/docker-general";
  import { faDocker } from "@fortawesome/free-brands-svg-icons";

  let pageId = $state("home");
  let alerts: AlertType[] = $state([]);

  function pushAlert(alert: AlertType) {
    alerts.push(alert);
  }

  function setAlerts(array: AlertType[]) {
    alerts = array;
  }

  function updatePage(newpageid: string) {
    pageId = newpageid;
  }

  // Docker Ready Check
  (async () => {
    const drs = await dockerReady();

    try {
      if (drs.Ready != true) {
        pushAlert({
          color: "red",
          content: drs.Msg,
          icon: faDocker,
        });

        console.error(drs);
      }
    } catch {
      pushAlert({
        color: "red",
        content:
          "Docker might not be available on your system, please restart Dampfer, check the web browser console or contact the developers!",
        icon: faDocker,
      });

      console.error(drs);
    }
  })();
</script>

<main>
  <LoginCheck {pushAlert} {updatePage} />

  <Sidebar
    className="sidebar"
    pageid={pageId}
    {updatePage}
    {alerts}
    {pushAlert}
    {setAlerts}
  />

  <Page className="page" pageid={pageId} {pushAlert} {updatePage} />
</main>
