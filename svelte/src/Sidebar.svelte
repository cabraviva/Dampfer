<script lang="ts">
  let { pageid, className, updatePage } = $props();
  import Fa from "svelte-fa";

  import "./app.scss";
  import "./sidebar.scss";
  import SidebarButton from "./sidebar/SidebarButton.svelte";
  import SidebarLogo from "./sidebar/SidebarLogo.svelte";
  import { faBox, faHouse } from "@fortawesome/free-solid-svg-icons";
</script>

<nav class={className + " sidebar-container"}>
  <SidebarLogo />
  <ul class="sidebar-list">
    <SidebarButton currentPageId={pageid} {updatePage} pageid="home">
      <Fa icon={faHouse} class="mr-3" /> Home
    </SidebarButton>
    <SidebarButton currentPageId={pageid} {updatePage} pageid="not-home">
      <Fa icon={faBox} class="mr-3" /> Not Home
    </SidebarButton>
  </ul>
</nav>
