<template>
  <div class="engine-container">
    <div class="engine-text">
      <el-table v-loading="engineLoading" align="center" :data="engineInfo">
        <el-table-column fixed prop="name" label="模型名" width="100px"></el-table-column>
        <el-table-column prop="params" label="参数" max-width="100px">
          <template slot-scope="scope">
            <div>{{ scope.row.params }}</div>
<!--            <div v-if="dataObject(scope.row.params)">-->
<!--              <pre style="overflow:auto;"><code>{{ JSON.stringify(scope.row.params, null, 4).replace(/\"/g, "")}}</code></pre>-->
<!--            </div>-->
<!--            <div v-else>{{ scope.row.params }}</div>-->
          </template>
        </el-table-column>
        <el-table-column prop="weight" label="权重" width="100px">
          <template slot-scope="scope">
            <div>{{ scope.row.weight }}</div>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script>
import { getEngineInfo } from '@/api/engine'

export default {
  data() {
    return {
      engineInfo: null,
      engineLoading: false
    }
  },
  created() {
    this.fetchEngineInfo()
  },
  methods: {
    fetchEngineInfo() {
      this.engineLoading = true
      getEngineInfo().then(rsp => {
        this.engineInfo = JSON.parse(rsp.data)["models"]
        this.engineInfo.forEach(function (item, index) {
          item.weight = item.params.weight
          delete item.params.weight
        })
        this.engineLoading = false
      })
    },
    dataObject(info) {
      try {
        return Object.prototype.toString.call(info) === "[object Object]";
      } catch (e) {
        return false;
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.engine {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>
