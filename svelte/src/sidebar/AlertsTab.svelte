<script lang="ts">
  let { alerts, pushAlert, setAlerts } = $props();

  import Fa from "svelte-fa";

  import "../app.scss";
  import "../sidebar.scss";
  import { faBell } from "@fortawesome/free-solid-svg-icons";

  import { Drawer, Toast, CloseButton, Button } from "flowbite-svelte";
  import { sineIn } from "svelte/easing";
  import type { AlertType } from "../types";
  import { faDocker } from "@fortawesome/free-brands-svg-icons";
  import { slide } from "svelte/transition";
  import { Indicator } from "flowbite-svelte";

  let hidden1 = $state(true);
  let transitionParams = {
    x: -320,
    duration: 200,
    easing: sineIn,
  };

  let newAlerts = $state(false);

  $effect(() => {
    let areThereNewAlerts = false;
    for (const alert of alerts) {
      if (alert.toastStatus !== false) {
        areThereNewAlerts = true;
      }
    }

    newAlerts = areThereNewAlerts;
  });
</script>

<!-- Alert Icon -->

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div onclick={() => (hidden1 = false)}>
  <Fa icon={faBell} />
  {#if newAlerts}
    <Indicator color="red" class="fixed" style="bottom: 4.4vh;left: 7.8vw;" />
  {/if}
</div>

<!-- Alerts Panel -->

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
      {#if alert.toastStatus !== false}
        <div class="mb-3">
          <Toast
            color={alert.color}
            transition={slide}
            dismissable={alert.canBeIgnored}
            bind:toastStatus={alert.toastStatus}
          >
            <svelte:fragment slot="icon">
              <div class="text-sm"><Fa icon={alert.icon} /></div>
            </svelte:fragment>

            <div>
              {alert.content}
            </div>

            {#if typeof alert.viewButton === "object"}
              <div class="mt-2">
                <Button on:click={alert.viewButton.onclick}>
                  {alert.viewButton.text || "View"}
                </Button>
              </div>
            {/if}
          </Toast>
        </div>
      {/if}
    {/each}
  </div>
</Drawer>
