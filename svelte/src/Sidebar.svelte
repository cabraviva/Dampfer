<script lang="ts">
  let { pageid, className, updatePage, alerts, pushAlert, setAlerts } =
    $props();
  import Fa from "svelte-fa";

  import "./app.scss";
  import "./sidebar.scss";
  import SidebarButton from "./sidebar/SidebarButton.svelte";
  import SidebarLogo from "./sidebar/SidebarLogo.svelte";
  import {
    faBox,
    faClockRotateLeft,
    faGears,
    faHouse,
    faSdCard,
    faServer,
    faUsers,
  } from "@fortawesome/free-solid-svg-icons";
  import { faDocker } from "@fortawesome/free-brands-svg-icons";
  import { whoami } from "./script/whoami";
  import SideBarTabs from "./sidebar/SideBarTabs.svelte";
  import AlertsTab from "./sidebar/AlertsTab.svelte";
  import type { AlertType } from "./types";
</script>

<nav class={className + " sidebar-container"}>
  <SidebarLogo />

  <ul class="sidebar-list">
    <SidebarButton currentPageId={pageid} {updatePage} pageid="home">
      <Fa icon={faHouse} class="mr-3" /> Home
    </SidebarButton>
    <SidebarButton currentPageId={pageid} {updatePage} pageid="containers">
      <Fa icon={faDocker} class="mr-3" /> Containers
    </SidebarButton>
    <SidebarButton currentPageId={pageid} {updatePage} pageid="images">
      <Fa icon={faBox} class="mr-3" /> Images
    </SidebarButton>
    <SidebarButton currentPageId={pageid} {updatePage} pageid="volumes">
      <Fa icon={faSdCard} class="mr-3" /> Volumes
    </SidebarButton>
    <SidebarButton currentPageId={pageid} {updatePage} pageid="system">
      <Fa icon={faServer} class="mr-3" /> System
    </SidebarButton>
    {#await whoami() then userInfo}
      {#if userInfo.admin}
        <SidebarButton currentPageId={pageid} {updatePage} pageid="backup">
          <Fa icon={faClockRotateLeft} class="mr-3" /> Backup
        </SidebarButton>
      {/if}

      {#if userInfo.systemAdmin}
        <SidebarButton currentPageId={pageid} {updatePage} pageid="users">
          <Fa icon={faUsers} class="mr-3" /> Users
        </SidebarButton>

        <SidebarButton currentPageId={pageid} {updatePage} pageid="settings">
          <Fa icon={faGears} class="mr-3" /> Settings
        </SidebarButton>
      {/if}
    {/await}
  </ul>

  <SideBarTabs {alerts} {pushAlert} {setAlerts} />
</nav>
