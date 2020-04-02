<template>
  <div class="user-container">
    <el-form ref="user" v-loading="userLoding" :model="user" label-width="100px">
      <el-form-item label="Nick Name">
        <el-input v-model="user.Name" />
      </el-form-item>
      <el-form-item label="User ID">
        <el-input v-model="user.UserId" disabled />
      </el-form-item>
      <el-form-item label="Gender">
        <el-select v-model="user.Gender" placeholder="Select Gender">
          <el-option v-for="item in genders" :key="item.value" :label="item.label" :value="item.value" />
        </el-select>
      </el-form-item>
      <el-form-item label="Age">
        <el-select v-model="user.Age" placeholder="Select Age">
          <el-option v-for="item in ages" :key="item.value" :label="item.label" :value="item.value" />
        </el-select>
      </el-form-item>
      <el-form-item label="Occupation">
        <el-select v-model="user.Occupation" placeholder="Select Occupation">
          <el-option v-for="item in occupations" :key="item.value" :label="item.label" :value="item.value" />
        </el-select>
      </el-form-item>
    </el-form>
    <div class="user-button-container">
      <el-button type="primary" icon="el-icon-upload" @click="UpdateUserInfo">Submit</el-button>
      <router-link to="/">
        <el-button type="danger" icon="el-icon-close">Close</el-button>
      </router-link>
    </div>
  </div>
</template>

<script>
import { getInfo, updateUserInfo } from '@/api/user'

export default {
  name: 'User',
  data() {
    return {
      user: {},
      userLoding: false,
      genders: [
        { value: 'F', label: 'Female' },
        { value: 'M', label: 'Male' }
      ],
      ages: [
        { value: 1, label: 'Under 18' },
        { value: 18, label: '18 - 24' },
        { value: 25, label: '25 - 34' },
        { value: 35, label: '35 - 44' },
        { value: 45, label: '45 - 49' },
        { value: 50, label: '50 - 55' },
        { value: 56, label: '56 +' }
      ],
      occupations: [
        { value: 0, label: 'other' },
        { value: 1, label: 'academic/educator' },
        { value: 2, label: 'artist' },
        { value: 3, label: 'clerical/admin' },
        { value: 4, label: 'college/grad student' },
        { value: 5, label: 'customer service' },
        { value: 6, label: 'doctor/health care' },
        { value: 7, label: 'executive/managerial' },
        { value: 8, label: 'farmer' },
        { value: 9, label: 'homemaker' },
        { value: 10, label: 'K-12 student' },
        { value: 11, label: 'lawyer' },
        { value: 12, label: 'programmer' },
        { value: 13, label: 'retired' },
        { value: 14, label: 'sales/marketing' },
        { value: 15, label: 'scientist' },
        { value: 16, label: 'self-employed' },
        { value: 17, label: 'technician/engineer' },
        { value: 18, label: 'tradesman/craftsman' },
        { value: 19, label: 'unemployed' },
        { value: 20, label: 'writer' }
      ]
    }
  },
  created() {
    this.fetchUserInfo()
  },
  methods: {
    UpdateUserInfo() {
      this.userLoding = true
      updateUserInfo(this.user).then(rsp => {
        this.userLoding = false
        this.$message({
          message: 'Submit Engine Information Success!',
          type: 'success'
        })
      })
    },
    fetchUserInfo() {
      this.userLoding = true
      getInfo().then(rsp => {
        this.user = rsp.data
        this.userLoding = false
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.user {
  &-container {
    margin: 30px;
    border-radius: 4px;
    min-height: 36px;
    padding: 20px;
    opacity:0.9;
    background: #d3dce6;
  }
  &-button-container {
    text-align: right;
  }
}
</style>
