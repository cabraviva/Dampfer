<script lang="ts">
  let { pageid, className, pushAlert, updatePage } = $props();

  import "./app.scss";
  import "./page.scss";
  import Backup from "./pages/backup/Backup.svelte";
  import Containers from "./pages/containers/Containers.svelte";
  import Home from "./pages/home/Home.svelte";
  import Settings from "./pages/settings/Settings.svelte";
  import System from "./pages/system/System.svelte";
  import Users from "./pages/users/Users.svelte";
  import Images from "./pages/images/Images.svelte";
  import Volumes from "./pages/volumes/Volumes.svelte";
</script>

<main class={className + " page-container"}>
  {#if pageid === "home"}
    <Home {pushAlert} {updatePage} />
  {:else if pageid === "containers"}
    <Containers {pushAlert} {updatePage} />
  {:else if pageid === "images"}
    <Images {pushAlert} {updatePage} />
  {:else if pageid === "volumes"}
    <Volumes {pushAlert} {updatePage} />
  {:else if pageid === "system"}
    <System {pushAlert} {updatePage} />
  {:else if pageid === "backup"}
    <Backup {pushAlert} {updatePage} />
  {:else if pageid === "users"}
    <Users {pushAlert} {updatePage} />
  {:else if pageid === "settings"}
    <Settings {pushAlert} {updatePage} />
  {:else}
    <h1>404: Page could not be found</h1>
  {/if}
</main>
