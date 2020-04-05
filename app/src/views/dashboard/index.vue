<template>
  <div class="dashboard-container">
    <el-row v-for="item in headerList" :key="item.key" v-loading="movieLoading" class="grid-content bg-purple" style="padding-bottom: 20px">
      <div>
        <span style="font-size: 20px;padding-top: 15px;">{{ item.name }}</span>
        <span style="float: right;color: #1482f0;padding-right: 5px;padding-top: 10px">
          <router-link :to="item.link">更多</router-link>
        </span>
      </div>
      <HR style="FILTER: alpha(opacity=100, finishopacity=0,style=1)" width="100%" color="#c0c0c0" size="3" />
      <div v-if="movieLists[item.key] != null">
        <div v-for="(m, index) in movieLists[item.key]" :key="index" style="padding:2px">
          {{ addMovie2TmpList(m) }}
          <el-row v-if="rowList.length === 5">
            <div v-for="(movie,idx) in rowList" :key="idx">
              <el-col :span="4" :offset="idx> 0 ? 1 : 0">
                <el-card>
                  <div class="block">
                    <el-image fit="fill" :src="movie.img">
                      <div slot="error" class="image-slot">
                        <i class="el-icon-picture-outline" />
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
    </el-row>
  </div>
</template>

<script>
import { getMovieLists } from '@/api/movies'

export default {
  name: 'Home',
  data() {
    return {
      headerList: [
        { name: '猜你喜欢', link: '/movie/recommend', key: 'movieRecommend' },
        // { name: '大家都在看', link: '/', key:'' },
        { name: '评分最高', link: '/movie/toprank', key: 'movieTopRank' },
        { name: '最近观看', link: '/movie/recently', key: 'movieRecently' }
      ],
      defaultReq: {
        movie_recently_num: 5,
        movie_top_rank_num: 10,
        movie_recommend_num: 0
      },
      movieLoading: false,
      movieLists: {
        'movieRecently': [],
        'movieRecommend': [],
        'movieTopRank': []
      },
      rowList: []
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
    num2Fixed(attr) {
      if (this.movieLists[attr] != null) {
        this.movieLists[attr].forEach(movie => {
          movie.avgrating = parseFloat(movie.avgrating.toFixed(1))
        })
      }
    },
    fetchMovieLists() {
      this.movieLoading = true
      getMovieLists(this.defaultReq).then(rsp => {
        this.movieLists = rsp.data
        this.headerList.forEach(item => {
          this.num2Fixed(item.key)
        })
        this.movieLoading = false
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.dashboard {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
.grid-content {
  border-radius: 4px;
  min-height: 36px;
  margin-bottom: 20px;
  padding: 5px;
  opacity:0.8;
}
.bg-purple {
  background: #d3dce6;
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
