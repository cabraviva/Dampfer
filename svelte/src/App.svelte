<script lang="ts">
  import "./app.css";
  import LoginCheck from "./popups/LoginCheck.svelte";
  import Sidebar from "./Sidebar.svelte";
  import Page from "./Page.svelte";

  let pageId = $state("home");

  function updatePage(newpageid: string) {
    pageId = newpageid;
  }
</script>

<main>
  <LoginCheck />

  <Sidebar className="sidebar" pageid={pageId} {updatePage} />

  <Page className="page" pageid={pageId} />
</main>
