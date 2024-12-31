<script lang="ts">
  let { children } = $props();

  import "../app.scss";
  import "../sidebar.scss";
</script>

<li class="sidebar-tab-element">
  {@render children?.()}
</li>
