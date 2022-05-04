<template>
  <div class="col-lg-12">
    <div class="iq-edit-list-data">
      <div class="card">
        <div class="card-header d-flex justify-content-between">
          <div class="header-title">
            <h4 class="card-title">Personal Information Settings</h4>
          </div>
        </div>
        <div class="card-body">
          <form>
            <div class="form-group row align-items-center text-center">
              <div class="col-md-12">
                <div class="profile-img-edit">
                  <img
                    class="profile-pic"
                    src="https://templates.iqonic.design/socialv/bs5/html/dist/assets/images/user/11.png"
                    alt="profile-pic"
                  />
                  <div class="p-image">
                    <i class="ri-pencil-line upload-button text-white"></i>
                    <input class="file-upload" type="file" accept="image/*" />
                  </div>
                </div>
              </div>
            </div>
            <div class="row align-items-center">
              <div v-for="input of inputs" :key="input.id" class="form-group" :class="input.colType">
                <div v-if="input.id !== 'aboutMe'">
                  <label for="fname" class="form-label">{{input.label}}</label>
                  <input
                    :type="input.type"
                    class="form-control"
                    :id="input.id"
                    :minlength="input.minLenght ? input.minLenght : _"
                    :placeholder="input.label"
                  />
                </div>
                <div v-else>
                  <label class="form-label">About me:</label>
                  <textarea
                    class="form-control"
                    name="aboutMe"
                    rows="5"
                    style="line-height: 22px"
                    placeholder="About me"
                  ></textarea>
                </div>
              </div>
            </div>
            <button type="submit" class="btn btn-primary me-2">Update</button>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ProfileSettings",
  data: () => ({
    inputs: [
      {
        label: "First Name",
        id: "firstName",
        type: "text",
        colType: "col-sm-6",
      },
      {
        label: "Last Name",
        id: "lastName",
        type: "text",
        colType: "col-sm-6",
      },
      {
        label: "Nickname",
        id: "nickname",
        type: "text",
        minLenght: 4,
        colType: "col-sm-6",
      },
      {
        label: "Date Of Birth",
        id: "birthDate",
        type: "date",
        colType: "col-sm-6",
      },
      {
        label: "Email",
        id: "email",
        type: "email",
        colType: "col-sm-6",
      },
      {
        label: "About me",
        id: "aboutMe",
        type: "text",
        colType: "col-sm-12",
      },                  
    ]
  })
};
</script>

<style>
</style>