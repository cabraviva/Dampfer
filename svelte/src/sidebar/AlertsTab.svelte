<script lang="ts">
  let { alerts, pushAlert } = $props();

  import Fa from "svelte-fa";

  import "../app.scss";
  import "../sidebar.scss";
  import { faBell } from "@fortawesome/free-solid-svg-icons";

  import { Drawer, Toast, CloseButton } from "flowbite-svelte";
  import { sineIn } from "svelte/easing";
  import type { AlertType } from "../types";
  import { faDochub, faDocker } from "@fortawesome/free-brands-svg-icons";
  let hidden1 = $state(true);
  let transitionParams = {
    x: -320,
    duration: 200,
    easing: sineIn,
  };

  // TODO: Remove this later
  setInterval(() => {
    pushAlert({
      icon: faDocker,

      color: "blue",

      content: "Docker Hub was not able to load!",

      viewButton: {
        text: "Fix!",
        onclick: () => alert("Was not able to fix the error!"),
      },

      canBeIgnored: false,
    } as AlertType);
  }, 5_000);
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div onclick={() => (hidden1 = false)}><Fa icon={faBell} /></div>

<Drawer
  transitionType="fly"
  {transitionParams}
  bind:hidden={hidden1}
  id="sidebar1"
>
  <div class="flex items-center">
    <h5
      id="drawer-label"
      class="inline-flex items-center mb-4 text-base font-semibold text-gray-500 dark:text-gray-400"
    >
      ALERTS
    </h5>
    <CloseButton
      on:click={() => (hidden1 = true)}
      class="mb-4 text-primary-600"
    />
  </div>

  <div class="flex flex-col justify-start">
    {#each alerts as alert}
      <div class="mb-3">
        <!-- TODO: Onclose, view button, etc... -->
        <Toast color={alert.color}>
          <svelte:fragment slot="icon">
            <div class="text-sm"><Fa icon={alert.icon} /></div>
          </svelte:fragment>
          {alert.content}
        </Toast>
      </div>
    {/each}
  </div>
</Drawer>
