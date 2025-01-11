<script lang="ts">
  import {
    faCirclePlus,
    faTrash,
    faUserPlus,
    faXmark,
  } from "@fortawesome/free-solid-svg-icons";
  import "../../app.scss";
  import UserManagementBox from "./UserManagementBox.svelte";
  import Fa from "svelte-fa";
  import { whoami } from "../../script/whoami";
  import {
    createUser,
    deleteUser,
    listUsers,
    type UserList,
  } from "../../script/user-api";
  import SmallPopup from "../../popups/SmallPopup.svelte";
  import { Button, Modal } from "flowbite-svelte";

  let { pushAlert, updatePage } = $props();

  let users: UserList = $state([]);

  function refetchUsers() {
    listUsers().then((userVal) => {
      users = userVal;
    });
  }

  refetchUsers();

  let showPopup = $state(false);

  import { Select, Label, Input, Helper } from "flowbite-svelte";
  let selectedPermission = $state("insight");
  let selectedUsername = $state("");
  let selectedPassword = $state("");
  let permissionCheckbox = $state([
    { value: "system-admin", name: "System Admin (Full privileges)" },
    { value: "admin", name: "Admin (Can manage services but not users)" },
    {
      value: "insight",
      name: "Insight (Lowest permission, can't make changes)",
    },
  ]);

  import { ExclamationCircleOutline } from "flowbite-svelte-icons";

  let changingUser = $state("");
  let mayShowAnyChangePopup = $derived(changingUser !== "");
  let showChangePermissionsPopup = $state(false);
  let showChangePasswordPopup = $state(false);
  let showDeleteModal = $state(false);
</script>

<main
  class="flex justify-start flex-col w-full h-full items-start pt-4 pb-4 pl-8 pr-8"
>
  <!-- Heading + Create User Button -->
  <h1 class="text-3xl mb-8 heading">
    <span class="l">Manage users</span>
    <!-- svelte-ignore a11y_click_events_have_key_events -->
    <!-- svelte-ignore a11y_no_static_element_interactions -->
    <span class="r" onclick={() => (showPopup = true)}>
      <Fa icon={faCirclePlus} />
    </span>
  </h1>

  <!-- Create User Popup -->
  <SmallPopup
    show={showPopup}
    onclose={() => (showPopup = false)}
    title="Create a new user"
  >
    <div class="pop-c">
      <div class="inputs">
        <div class="mb-6">
          <Label for="username" class="mb-2">Username</Label>
          <Input
            type="text"
            id="username"
            placeholder="user"
            required
            bind:value={selectedUsername}
          />
        </div>

        <div class="mb-6">
          <Label class="mb-2" for="password">Password</Label>
          <Input
            type="password"
            id="password"
            placeholder="•••••••••"
            bind:value={selectedPassword}
            required
          />
        </div>

        <Label>
          Permission
          <Select
            class="mt-2"
            items={permissionCheckbox}
            bind:value={selectedPermission}
          />
        </Label>
      </div>
      <div class="buttons">
        <Button color="alternative" on:click={() => (showPopup = false)}
          >Cancel</Button
        >
        <Button
          on:click={async () => {
            showPopup = false;

            const success = await createUser({
              username: selectedUsername,
              password: selectedPassword,
              permission: selectedPermission as
                | "insight"
                | "admin"
                | "system-admin",
            });

            refetchUsers();

            if (success) {
              pushAlert({
                icon: faUserPlus,
                color: "green",
                content: "User was created successfully",
              });
            } else {
              pushAlert({
                icon: faUserPlus,
                color: "red",
                content:
                  "Error while trying to create user. You might want to check the logs!",
              });
            }
          }}>Save</Button
        >
      </div>
    </div>
  </SmallPopup>

  <!-- Change Permissions Popup -->
  <SmallPopup
    title={"Change permission for: " + changingUser}
    show={mayShowAnyChangePopup && showChangePermissionsPopup}
    onclose={() => (showChangePermissionsPopup = false)}
  >
    XYZ TODO:
  </SmallPopup>

  <!-- Confirm delete popup -->
  <Modal bind:open={showDeleteModal} size="xs" autoclose>
    <div class="text-center">
      <ExclamationCircleOutline
        class="mx-auto mb-4 text-gray-400 w-12 h-12 dark:text-gray-200"
      />
      <h3 class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
        Are you sure you want to delete the user {changingUser}?
      </h3>
      <Button
        color="primary"
        class="me-2"
        on:click={async () => {
          const success = await deleteUser(changingUser);

          refetchUsers();

          if (success) {
            pushAlert({
              icon: faTrash,
              color: "green",
              content: `User ${changingUser} was deleted successfully`,
            });
          } else {
            pushAlert({
              icon: faXmark,
              color: "red",
              content: `Error while trying to delete user ${changingUser}. You might want to check the logs!`,
            });
          }
        }}>Yes, I'm sure</Button
      >
      <Button color="alternative">No, cancel</Button>
    </div>
  </Modal>

  <!-- User Management Boxes -->
  {#await whoami() then userInfo}
    {#each users as user}
      <UserManagementBox
        {pushAlert}
        {updatePage}
        username={user.username}
        permission={user.permission}
        isme={userInfo.username === user.username}
        onChangePermission={() => {
          changingUser = user.username;
          showChangePermissionsPopup = true;
        }}
        onChangePassword={() => {
          changingUser = user.username;
          showChangePasswordPopup = true;
        }}
        onDelete={() => {
          changingUser = user.username;
          showDeleteModal = true;
        }}
      />
    {/each}
  {/await}
</main>

<!-- Local styling -->
<style lang="scss">
  .pop-c {
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: space-between;
    align-items: center;
    flex-direction: column;

    padding-top: 1rem;
    padding-bottom: 1rem;

    .buttons {
      width: 100%;
      display: flex;
      flex-direction: row;
      justify-content: space-around;
      align-items: center;
    }
  }

  .heading {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    width: 100%;

    .r {
      color: #102e4a;
      &:hover {
        cursor: pointer;
        color: #0d7e2b;
      }
    }
  }
</style>
