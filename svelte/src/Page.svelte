<script lang="ts">
  let { pageid, className } = $props();

  import "./app.scss";
  import "./page.scss";
  import Backup from "./pages/backup/Backup.svelte";
  import Containers from "./pages/containers/Containers.svelte";
  import Home from "./pages/home/Home.svelte";
  import Settings from "./pages/settings/Settings.svelte";
  import System from "./pages/system/System.svelte";
  import Users from "./pages/users/Users.svelte";
  import Images from "./pages/containers/images/Images.svelte";
  import Volumes from "./pages/volumes/Volumes.svelte";
</script>

<main class={className + " page-container"}>
  {#if pageid === "home"}
    <Home />
  {:else if pageid === "containers"}
    <Containers />
  {:else if pageid === "images"}
    <Images />
  {:else if pageid === "volumes"}
    <Volumes />
  {:else if pageid === "system"}
    <System />
  {:else if pageid === "backup"}
    <Backup />
  {:else if pageid === "users"}
    <Users />
  {:else if pageid === "settings"}
    <Settings />
  {:else}
    <h1>404: Page could not be found</h1>
  {/if}
</main>
