<script lang="ts">
  import "../../app.scss";
  import "./settingsPage.scss";
  import logo from "../../../public/logo.svg";

  let { pushAlert, updatePage } = $props();

  import { Tabs, TabItem, Label, Input, Button } from "flowbite-svelte";
  import SettingsPage from "./SettingsPage.svelte";
  import { whoami } from "../../script/whoami";
  import {
    faArrowRight,
    faCheck,
    faXmark,
  } from "@fortawesome/free-solid-svg-icons";
  import Fa from "svelte-fa";
  import { changeMyPassword } from "../../script/user-api";

  let oldPassword = $state("");
  let newPassword = $state("");
  let newPassword2 = $state("");
</script>

<main class="w-full h-full">
  <!-- TODO: Start settings page, User tab, Be able to change password -->

  <div class="tabs">
    <Tabs>
      <TabItem open title="User">
        <SettingsPage>
          {#await whoami() then userInfo}
            <h1 class="text-2xl mb-4">Welcome, {userInfo.username}!</h1>
          {/await}

          <hr class="mt-4 mb-4" />

          <h2 class="text-l mb-3">Change password:</h2>
          <div>
            <div class="mb-2">
              <Label for="oldpassword" class="mb-2">Old password</Label>
              <Input
                type="password"
                id="oldpassword"
                placeholder="•••••••••"
                bind:value={oldPassword}
              />
            </div>

            <div class="mb-2">
              <Label for="newpassword" class="mb-2">New password</Label>
              <Input
                type="password"
                id="newpassword"
                placeholder="•••••••••"
                bind:value={newPassword}
              />
            </div>

            <div class="mb-2">
              <Label for="newpassword2" class="mb-2">Confirm new password</Label
              >
              <Input
                type="password"
                id="newpassword2"
                placeholder="•••••••••"
                bind:value={newPassword2}
              />
            </div>

            <Button
              class="mt-2 btn-submit"
              on:click={async () => {
                if (oldPassword.length === 0) {
                  return pushAlert({
                    icon: faXmark,
                    color: "red",
                    content:
                      "Was not able to change password! Old password must not be empty!",
                  });
                }

                if (newPassword.length < 8) {
                  return pushAlert({
                    icon: faXmark,
                    color: "red",
                    content:
                      "Was not able to change password! New password must be at least 8 characters long!",
                  });
                }

                if (newPassword !== newPassword2) {
                  return pushAlert({
                    icon: faXmark,
                    color: "red",
                    content:
                      "Was not able to change password! New password must be identical to confirm new password!",
                  });
                }

                const result = await changeMyPassword(oldPassword, newPassword);

                if (result === true) {
                  pushAlert({
                    icon: faCheck,
                    color: "green",
                    content: "Password was successfully changed!",
                  });
                } else {
                  return pushAlert({
                    icon: faXmark,
                    color: "red",
                    content: `Was not able to change password! ${result}`,
                  });
                }
              }}>Submit <Fa class="ml-2" icon={faArrowRight} /></Button
            >
          </div>
        </SettingsPage>
      </TabItem>
    </Tabs>
  </div>
</main>
