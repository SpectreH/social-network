<template>
  <div class="card">
    <div class="card-body">
        <div class="row">
          <div class="col-md-12 ps-4">
              <div class="tab-content">
                <div>
                  <div v-for="section of sections" :key="section.id">
                    <h4 class="mt-3">{{section.header}}</h4>
                    <hr>
                    <div v-for="element of section.content" :key="element.id"  class="row">
                      <div v-if="section.id !== 'about'" class="row">
                        <div class="col-3">
                            <h6>{{element.key}}</h6>
                        </div>
                        <div class="col-9">
                            <p v-if="element.value">{{element.value}}</p>
                            <p v-else style="font-style: italic;">Unknown</p>  
                        </div>
                      </div>
                      <div v-else class="row">
                        <div class="col-12">
                          <p v-if="element.value">{{element.value}}</p>
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
  data: () => ({
    sections: [
      {
        header: "Personal Information",
        id: "personal",
        content: [
          {
            id: "firstName",
            key: "First Name",
            value: "Denni"
          },
           {
            id: "lastName",
            key: "Last Name",
            value: "Karin"
          },
          {
            id: "birthDate",
            key: "Birth Date",
            value: "28.03.2002"
          },  
          {
            id: "nickname",
            key: "Nickname",
            value: ""
          },                          
        ],
      },
      {
        header: "Contact Information",
        id: "contact",
        content: [
          {
            id: "email",
            key: "Email",
            value: "test@gmail.com"
          }                          
        ],
      },
      {
        header: "About me",
        id: "about",
        content: [
          {
            id: "aboutMe",
            value: ""
          }                          
        ],
      }
    ]
  })
}
</script>


<style scoped>
p {
  margin: 0;
}
</style>