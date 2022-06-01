<template>
  <div class="col-12 iq-follow-request">
    <div
      class="
        iq-sub-card iq-sub-card-big
        d-flex
        align-items-center
        justify-content-between
      "
    >
      <div class="col-6 d-flex align-items-center">
        <img :src="request.avatar" class="avatar-40 rounded" alt="" />
        <div class="ms-3">
          <h6 class="mb-0">{{ request.author }}</h6>
          <p class="mb-0">{{ request.sub }}</p>
        </div>
      </div>
      <div class="col-6 d-flex align-items-center justify-content-end">
        <a
          href="javascript:void();"
          class="me-3 btn btn-primary rounded"
          >Confirm</a
        >
        <a
          href="javascript:void();"
          class="me-3 btn btn-secondary rounded"
          >Decline</a
        >
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "FollowRequest",
  props: { 
    request: {type: Object, default: () => {
      return {}
    }}
  }
}
</script>

<style>

</style>