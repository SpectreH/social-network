<template>
  <div class="d-flex align-items-center justify-content-between mb-3 flex-wrap">
    <div class="group-info d-flex align-items-center">
      <div class="me-3">
        <img class="rounded-circle img-fluid avatar-100" src="https://templates.iqonic.design/socialv/bs5/html/dist/assets/images/page-img/gi-1.jpg" alt="">
      </div>
      <div class="info">
        <h4>Developer's...</h4>
        <p class="mb-0"><i class="ri-lock-fill pe-2"></i>Private Group. 323 members</p>
      </div>
    </div>
    <div class="group-member d-flex align-items-center gap-2  mt-md-0 mt-2">
      <button type="submit" class="btn btn-primary mb-2">Invite</button>
      <button type="submit" class="btn btn-primary mb-2">Request to join</button> 
      <button type="submit" class="btn btn-primary mb-2">Follow</button> 
      <button type="submit" class="btn btn-primary mb-2">Unfollow</button>       
    </div>
  </div>
</template>

<script>
export default {
  name: "GroupHeader"
}
</script>

<style>

</style>