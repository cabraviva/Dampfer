<script lang="ts">
  import "./settingsPage.scss";
  let { children } = $props();
</script>

<div class="settings-page">
  {@render children?.()}
</div>
