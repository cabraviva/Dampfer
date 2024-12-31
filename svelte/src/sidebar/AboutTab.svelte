<script lang="ts">
  import Fa from "svelte-fa";

  import "../app.scss";
  import "../sidebar.scss";
  import { faInfoCircle } from "@fortawesome/free-solid-svg-icons";
  import SmallPopup from "../popups/SmallPopup.svelte";
  import { A } from "flowbite-svelte";

  let show = $state(false);

  function openAboutTab() {
    show = true;
    console.log("HI!");
  }

  function onClose() {
    show = false;
  }
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div onclick={openAboutTab}><Fa icon={faInfoCircle} /></div>

<SmallPopup {show} onclose={onClose} title="About Dampfer">
  <h1 class="text-primary-900 text-xl">
    Dampfer is a FOSS container-management tool
  </h1>
  <br />

  <p class="text-justify">
    It is largely based on docker compose and meant to be a perfect fit for home
    users and small projects.
    <br /><br />
    It was created by <A
      class="hover:underline"
      href="https://github.com/cabraviva">CabraViva</A
    >. To learn more visit the <A
      class="hover:underline"
      href="https://github.com/cabraviva/Dampfer">GitHub page</A
    >.
  </p>
</SmallPopup>
