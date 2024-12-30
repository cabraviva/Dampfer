<script lang="ts">
  let { currentPageId, pageid, children, updatePage } = $props();
  let active = $derived(currentPageId === pageid);

  import "../app.scss";
  import "../sidebar.scss";

  function clickAction() {
    updatePage(pageid);
  }
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
<li class={{ active, "sidebar-btn": true }} onclick={clickAction}>
  <button onclick={clickAction}>{@render children?.()}</button>
</li>
