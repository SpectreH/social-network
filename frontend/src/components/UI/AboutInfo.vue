<template>
  <div class="card">
    <div class="card-body">
        <div class="row">
          <div class="col-md-12 ps-4">
              <div class="tab-content">
                <div>
                  <div>
                    <h4 class="mt-3">Personal Information</h4>
                    <hr>
                    <div class="row">
                      <div class="row">
                        <div class="col-3">
                            <h6>First Name</h6>
                        </div>
                        <div class="col-9">
                            <p>{{ info.firstName }}</p>
                        </div>
                      </div>

                      <div class="row">
                        <div class="col-3">
                            <h6>Last Name</h6>
                        </div>
                        <div class="col-9">
                            <p>{{ info.lastName }}</p>
                        </div>
                      </div>

                      <div class="row">
                        <div class="col-3">
                            <h6>Birth Date</h6>
                        </div>
                        <div class="col-9">
                            <p>{{ info.birthDate }}</p>
                        </div>
                      </div>

                      <div class="row">
                        <div class="col-3">
                            <h6>Nickname</h6>
                        </div>
                        <div class="col-9">
                            <p v-if="info.nickname">{{ info.nickname }}</p>
                            <p v-else style="font-style: italic;">Unknown</p>  
                        </div>
                      </div>
                    </div>
                  </div>

                  <div>
                    <h4 class="mt-3">Contact Information</h4>
                    <hr>
                    <div class="row">
                      <div class="col-3">
                          <h6>Email</h6>
                      </div>
                      <div class="col-9">
                          <p>{{ info.email }}</p>
                      </div>
                    </div>
                  </div>

                  <div>
                    <h4 class="mt-3">About me</h4>
                    <hr>
                    <div class="row">
                      <div class="row">
                        <div class="col-12">
                          <p v-if="info.aboutMe">{{ info.aboutMe }}</p>
                          <p v-else style="font-style: italic;">Unknown</p>  
                        </div>
                      </div>
                    </div>
                  </div>

                </div>
              </div>
          </div>
        </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'AboutInfo',
  props: {
    info: {type: Object}
  },
  data: () => ({})
}
</script>


<style scoped>
p {
  margin: 0;
}
</style>