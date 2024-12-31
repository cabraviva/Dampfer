<script lang="ts">
  import Fa from "svelte-fa";

  import "../app.scss";
  import "../sidebar.scss";
  import { faUserCircle } from "@fortawesome/free-solid-svg-icons";

  function openProfileTab() {}
</script>

<Fa icon={faUserCircle} on:click={openProfileTab} />
