<script lang="ts">
  import { faCirclePlus } from "@fortawesome/free-solid-svg-icons";
  import "../../app.scss";
  import UserManagementBox from "./UserManagementBox.svelte";
  import Fa from "svelte-fa";
  import { whoami } from "../../script/whoami";
  import { listUsers } from "../../script/user-api";

  let { pushAlert, updatePage } = $props();
</script>

<main
  class="flex justify-start flex-col w-full h-full items-start pt-4 pb-4 pl-8 pr-8"
>
  <h1 class="text-3xl mb-8 heading">
    <span class="l">Manage users</span>
    <span class="r">
      <Fa icon={faCirclePlus} />
    </span>
  </h1>

  {#await whoami() then userInfo}
    {#await listUsers() then users}
      {#each users as user}
        <UserManagementBox
          {pushAlert}
          {updatePage}
          username={user.username}
          permission={user.permission}
          isme={userInfo.username === user.username}
        />
      {/each}
    {/await}
  {/await}
</main>

<style>
  .heading {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    width: 100%;
  }

  .heading .r {
    color: #102e4a;
  }

  .heading .r:hover {
    cursor: pointer;
    color: #0d7e2b;
  }
</style>
