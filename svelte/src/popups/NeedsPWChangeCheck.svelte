<script lang="ts">
  import { faKey } from "@fortawesome/free-solid-svg-icons";
  import "../app.scss";
  import { pwChangeCount } from "../script/whoami";
  const { pushAlert, updatePage } = $props();

  (async () => {
    const count = await pwChangeCount();

    if (count < 1) {
      pushAlert({
        color: "red",
        content: "Your password was never changed. Please set a new one now!",
        icon: faKey,
        viewButton: {
          text: "Settings",
          onclick: () => {
            updatePage("settings");
          },
        },
      });
    }
  })();
</script>
