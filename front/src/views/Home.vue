<template>
  <v-app app>
    <TopBar></TopBar>
<!-- blackground:../assets/touxiang.jpg -->
    <v-main class="blue lighten-3" id="bg">
      <v-container>
        <v-row>
          <v-col cols="12" md="12" >
            <v-sheet min-height="80vh" rounded="lg">
              <router-view :key="$route.path"></router-view>
            </v-sheet>
          </v-col>
        </v-row>
      </v-container>
    </v-main>

    <Footer></Footer>
  </v-app>
</template>

<script>
import TopBar from '../components/TopBar'
import Footer from '../components/Footer'
export default {
  components: { TopBar, Footer }
}
</script>
<style scoped>
#bg{
  /* background-image: url('../assets/cat.jpg'); */
    background-repeat: no-repeat;
    background-attachment: fixed;
    background-position: center;
}
</style>
