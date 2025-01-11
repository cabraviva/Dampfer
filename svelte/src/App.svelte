<script lang="ts">
  import "./app.css";
  import LoginCheck from "./popups/LoginCheck.svelte";
  import Sidebar from "./Sidebar.svelte";
  import Page from "./Page.svelte";
  import type { AlertType } from "./types";

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
</script>

<main>
  <LoginCheck />

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
