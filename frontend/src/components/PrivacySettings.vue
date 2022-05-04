<template>
  <div class="col-lg-12">
    <div class="card">
      <div class="card-header d-flex justify-content-between">
        <div class="header-title">
          <h4 class="card-title">Privacy Setting</h4>
        </div>
      </div>
      <div class="card-body">
        <div class="acc-privacy">
          <div class="data-privacy">
            <h4 class="mb-2">Account Privacy</h4>
            <div class="form-check form-check-inline">
              <input
                type="checkbox"
                class="form-check-input"
                id="acc-private"/>
              <label
                class="form-check-label privacy-status mb-2"
                for="acc-private"
                >Private Account</label>
            </div>
            <p>
              Lorem Ipsum is simply dummy text of the printing and typesetting
              industry. Lorem Ipsum has been the industry's standard dummy text
              ever since the 1500s, when an unknown printer took a galley of
              type and scrambled it to make a type specimen book
            </p>
          </div>
          <button type="submit" class="btn btn-primary me-2">Update</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "PrivacySettings",
};
</script>

<style>
</style>