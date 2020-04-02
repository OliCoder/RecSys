<template>
  <div class="engine-container">
    <div class="engine-text">
      <el-table border highlight-current-row align="center" :data="engineInfo['models']" style="border-radius: 10px;">
        <el-table-column type="index" label="" />
        <el-table-column prop="name" label="Model Name" width="200px">
          <template slot-scope="scope">
            <el-select v-model="scope.row.name" placeholder="Model">
              <el-option v-for="item in algType" :key="item" :label="item" :value="item" />
            </el-select>
          </template>
        </el-table-column>
        <el-table-column prop="params" label="Params">
          <template slot-scope="scope">
            <el-input v-model="scope.row.params" />
          </template>
        </el-table-column>
        <el-table-column prop="weight" label="Weight" width="201px">
          <template slot-scope="scope">
            <el-input-number v-model="scope.row.weight" />
          </template>
        </el-table-column>
        <el-table-column prop="inuse" label="Enable" width="75px">
          <template slot-scope="scope">
            <el-switch v-model="scope.row.inuse" active-color="#13ce66" inactive-color="#ff4949" />
          </template>
        </el-table-column>
      </el-table>
      <div class="engine-button-container">
        <el-button icon="el-icon-plus" @click="addList" />
        <el-button type="primary" icon="el-icon-upload" @click="UpDateEngineInfo">Submit</el-button>
        <router-link to="/">
          <el-button type="danger" icon="el-icon-close">Close</el-button>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
import { getEngineInfo, UpdateEngineGroups } from '@/api/engine'

export default {
  data() {
    return {
      algType: [
        'CONTENT',
        'POPU',
        'ALS'
      ],
      engineInfo: { 'model': [] },
      engineLoading: false,
      submitText: 'Submit',
      submitLoading: false
    }
  },
  created() {
    this.fetchEngineInfo()
  },
  methods: {
    addList() {
      var sum = 0
      this.engineInfo['models'].forEach(item => {
        sum += item.weight
      })
      this.engineInfo['models'].unshift({ weight: 1 - sum })
    },
    fetchEngineInfo() {
      this.engineLoading = true
      getEngineInfo().then(rsp => {
        this.engineInfo = JSON.parse(rsp.data)
        this.engineInfo['models'].forEach(function(item, index) {
          item.weight = item.params.weight
          delete item.params.weight
          item.params = JSON.stringify(item.params)
          item.inuse = true
        })
        this.engineLoading = false
      })
    },
    UpDateEngineInfo() {
      this.submitLoading = true
      this.submitText = 'Loading'
      this.engineInfo['models'] = this.engineInfo['models'].filter(item => item.inuse === true)

      const engineInfo = JSON.parse(JSON.stringify(this.engineInfo))
      engineInfo['models'].forEach(function(item, index) {
        item.params = JSON.parse(item.params)
        item.params.weight = item.weight
        delete item.weight
        delete item.inuse
      })
      UpdateEngineGroups(engineInfo).then(rsp => {
        this.submitText = 'Submit'
        this.submitLoading = false
        this.$message({
          message: 'Submit Engine Information Success!',
          type: 'success'
        })
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.engine {
  &-container {
    margin: 10px;
    border-radius: 4px;
    min-height: 36px;
    padding: 20px;
    opacity:0.9;
    background: #d3dce6;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
  &-button-container {
    text-align: right;
  }
}
</style>
