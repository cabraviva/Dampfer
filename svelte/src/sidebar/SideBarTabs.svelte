<script lang="ts">
  let { alerts, pushAlert, setAlerts } = $props();

  import "../app.scss";
  import "../sidebar.scss";
  import SideBarTabElement from "./SideBarTabElement.svelte";
  import ProfileTab from "./ProfileTab.svelte";
  import AlertsTab from "./AlertsTab.svelte";
  import AboutTab from "./AboutTab.svelte";
</script>

<ul class="sidebar-tabs">
  <SideBarTabElement>
    <ProfileTab />
  </SideBarTabElement>
  <SideBarTabElement>
    <AlertsTab {alerts} {pushAlert} {setAlerts} />
  </SideBarTabElement>
  <SideBarTabElement>
    <AboutTab />
  </SideBarTabElement>
</ul>
