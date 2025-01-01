<script lang="ts">
  import Fa from "svelte-fa";
  import "../../app.scss";
  import "./SuggestionItem.scss";

  let { icon, href, text } = $props();
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
  onclick={href}
  class="suggestion-item flex flex-row justify-start items-center"
>
  <div class="sugg-icon">
    <Fa {icon} />
  </div>

  <div class="sugg-text">
    {text}
  </div>
</div>
