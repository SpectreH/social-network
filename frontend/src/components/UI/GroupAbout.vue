<template>
  <div class="card">
    <div class="card-header d-flex justify-content-between">
      <div class="header-title">
        <h4 class="card-title">About</h4>
      </div>
    </div>
    <div class="card-body">
      <ul class="list-inline p-0 m-0">
        <li>
          <h6>Description:</h6>
          <p class="mb-0">{{ group.description }}</p>
        </li>
        <hr class="m-2 mx-0 me-0">
        <li>
          <div class="d-flex gap-2 align-items-center">
            <h6>Status: </h6>
            <div class="d-flex gap-1 align-items-center">
              <div class="flex-shrink-0">
                <i v-if="group.private" class="ri-lock-fill h4"></i>
              </div>
              <div class="flex-grow-1">
                <h6>{{ group.private ? "Private" : "Public"}}</h6>
              </div>
            </div>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  name: "GroupAbout",
  props: {
    group: {type: Object}
  },
}
</script>

<style>

</style>