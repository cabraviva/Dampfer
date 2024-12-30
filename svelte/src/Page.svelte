<script lang="ts">
  let { pageid, className } = $props();

  import "./app.scss";
  import "./page.scss";
</script>

<main class={className + " page-container"}>This is the page</main>
