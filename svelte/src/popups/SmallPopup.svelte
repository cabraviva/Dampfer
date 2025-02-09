<script lang="ts">
  import Fa from "svelte-fa";
  import "./SmallPopup.scss";
  import { faXmark } from "@fortawesome/free-solid-svg-icons";

  let { children, show, onclose, title } = $props();
</script>

{#if show}
  <div class="small-popup">
    <div class="small-popup-top-bar">
      <h2>{title}</h2>

      <!-- svelte-ignore a11y_click_events_have_key_events -->
      <!-- svelte-ignore a11y_no_static_element_interactions -->
      <div onclick={onclose}><Fa icon={faXmark} /></div>
    </div>
    <div class="small-popup-content" style="overflow: auto;">
      {@render children?.()}
    </div>
  </div>
{/if}
