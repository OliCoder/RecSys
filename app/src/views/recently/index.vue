<template>
  <el-container v-loading="movieLoading">
    <el-main>
      <div class="grid-content bg-purple">
        <div v-for="(item, index) in movieLists" :key="index" style="padding:2px">
          {{ addMovie2TmpList(item) }}
          <el-row v-if="rowList.length === 5">
            <div v-for="(movie,idx) in rowList" :key="idx">
              <el-col :span="4" :offset="idx> 0 ? 1 : 0">
                <el-card>
                  <div class="block">
                    <el-image fit="fill" :src="movie.img">
                      <div slot="error" class="image-slot">
                        <el-image fit="fill" :src="errImg" />
                      </div>
                    </el-image>
                  </div>
                  <div class="movie">
                    <div class="movie-title" :title="movie.title">{{ movie.title }}</div>
                    <div class="movie-genre" :title="movie.genre">{{ movie.genre }}</div>
                    <el-rate v-model="movie.avgrating" disabled show-score text-color="#ff9900" score-template="{value}" />
                  </div>
                </el-card>
              </el-col>
            </div>
            {{ clearTmpList() }}
          </el-row>
        </div>
      </div>
    </el-main>
  </el-container>
</template>

<script>
import { getMovieLists } from '@/api/movies'
import ErrImg from '@/assets/tmp.png'

export default {
  name: 'Recently',
  data() {
    return {
      defaultReq: {
        movie_recently_num: 50
      },
      movieLoading: false,
      movieLists: [],
      rowList: [],
      errImg: ErrImg
    }
  },
  created() {
    this.fetchMovieLists()
  },
  methods: {
    clearTmpList() {
      delete this.rowList
      this.rowList = []
    },
    addMovie2TmpList(movie) {
      this.rowList.push(movie)
    },
    num2Fixed() {
      if (this.movieLists != null) {
        this.movieLists.forEach(movie => {
          movie.avgrating = parseFloat(movie.avgrating.toFixed(1))
        })
      }
    },
    fetchMovieLists() {
      this.movieLoading = true
      getMovieLists(this.defaultReq).then(rsp => {
        this.movieLists = rsp.data['movieRecently']
        this.num2Fixed()
        this.movieLoading = false
      })
    }
  }
}
</script>

<style lang="scss" scoped>
  .grid-content {
    border-radius: 4px;
    min-height: 36px;
    margin-bottom: 20px;
    padding: 10px;
    opacity:0.8;
  }
  .bg-purple {
    background: #d3dce6;
  }
  .block {
    /*width: 82px;*/
    /*height: 120px;*/
    /*margin-left: auto;*/
    /*margin-right: auto;*/
  }
  .movie {
  &-title {
     max-width: 200px;
     text-overflow: ellipsis;
     overflow: hidden;
     white-space: nowrap;
     padding: 3px;
     font-size: 16px
   }
  &-title:hover {
     max-width: 200px;
     text-overflow: ellipsis;
     overflow: hidden;
     white-space: nowrap;
     padding: 3px;
     font-size: 16px;
     text-decoration:underline;
     cursor:pointer;
   }
  &-genre {
     max-width: 200px;
     text-overflow: ellipsis;
     overflow: hidden;
     white-space: nowrap;
     font-size: 13px;
     color: #999;
   }
  &-genre:hover {
     max-width: 200px;
     text-overflow: ellipsis;
     overflow: hidden;
     white-space: nowrap;
     font-size: 13px;
     color: #999;
     text-decoration:underline;
     cursor:pointer;
   }
  }
</style>
