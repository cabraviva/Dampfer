<script lang="ts">
  import Fa from "svelte-fa";

  import "../app.scss";
  import "../sidebar.scss";
  import { faBell } from "@fortawesome/free-solid-svg-icons";
  import SmallPopup from "../popups/SmallPopup.svelte";

  function openAlertsTab() {}
</script>

<Fa icon={faBell} on:click={openAlertsTab} />
