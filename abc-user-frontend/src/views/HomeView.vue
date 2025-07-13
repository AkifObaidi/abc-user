<script setup>
import './HomeView.css'
import { ref, onMounted } from 'vue'
import userApi from '@/api/users'
import { ElMessageBox, ElMessage } from 'element-plus'
import UserForm from '@/components/UserForm.vue'
import { useRouter } from 'vue-router'
import { View, Edit, Delete } from '@element-plus/icons-vue';

const users = ref([])
const loading = ref(false)
const error = ref(null)
const total = ref(0)
const page = ref(1)
const pageSize = ref(15)

const showEditModal = ref(false)
const editingUser = ref(null)
const formLoading = ref(false)

const router = useRouter()

const loadUsers = async () => {
  loading.value = true
  error.value = null
  try {
    const response = await userApi.getUsers({ page: page.value, limit: pageSize.value })
    users.value = response.data.data
    total.value = response.data.total || response.data.data.length
  } catch (err) {
    error.value = 'Failed to load users.'
    ElMessage.error('Failed to load users.')
    console.error(err)
  } finally {
    loading.value = false
  }
}

const handleDelete = async (user) => {
  try {
    await ElMessageBox.confirm(`Delete user "${user.name}"?`, 'Confirm Delete', {
      type: 'warning',
      confirmButtonText: 'Delete',
      cancelButtonText: 'Cancel',
    })
    await userApi.deleteUser(user.id)
    ElMessage.success('User deleted')
    loadUsers()
  } catch (err) {
    if (err !== 'cancel') {
      ElMessage.error('Failed to delete user.')
    }
  }
}

const openEditModal = (user) => {
  editingUser.value = { ...user } // shallow clone
  showEditModal.value = true
}

const closeEditModal = () => {
  showEditModal.value = false
  editingUser.value = null
}

const submitEdit = async (updatedUser) => {
  formLoading.value = true
  try {
    await userApi.updateUser(editingUser.value.id, updatedUser)
    ElMessage.success('User updated successfully')
    closeEditModal()
    loadUsers()
  } catch (err) {
    ElMessage.error('Failed to update user')
    console.error(err)
  } finally {
    formLoading.value = false
  }
}

const handleAdd = () => {
  router.push('/new');
}

onMounted(loadUsers)
</script>

<template>
  <div class="home">
    <div class="header">
      <h1>ABC User Management</h1>
      <el-button type="primary" @click="handleAdd">Add User</el-button>
    </div>

    <el-alert
      v-if="error"
      title="Error"
      type="error"
      :closable="false"
      style="margin-bottom: 1rem;"
    />

    <el-card v-loading="loading" class="table-card" shadow="hover">
      <el-table :data="users" style="width: 100%">
        <el-table-column prop="name" label="Name" />
        <el-table-column prop="email" label="Email" />
        <el-table-column prop="age" label="Age" width="80" />
        <el-table-column
          prop="createdAt"
          label="Created"
          width="210"
          :formatter="row => new Date(row.createdAt).toLocaleString()"
        />
        <el-table-column
          prop="updatedAt"
          label="Updated"
          width="210"
          :formatter="row => new Date(row.updatedAt).toLocaleString()"
        />
        <el-table-column label="Actions" width="180">
          <template #default="{ row }">
            <el-button type="info" :icon="View" circle @click="() => router.push(`/user/${row.id}`)"></el-button> 
            <el-button type="primary" :icon="Edit" circle @click="() => openEditModal(row)"></el-button> 
            <el-button type="danger" :icon="Delete" circle @click="() => handleDelete(row)"></el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="table-footer">
        <div class="user-count">
          Showing {{ users.length }} of {{ total }} users
        </div>

        <el-pagination
          background
          layout="prev, pager, next"
          :total="total"
          :page-size="pageSize"
          :current-page="page"
          @current-change="newPage => { page = newPage; loadUsers() }"
        />
      </div>
    </el-card>

    <el-dialog
      v-model="showEditModal"
      title="Edit User"
      width="400px"
      :before-close="closeEditModal"
    >
      <UserForm
        v-if="editingUser"
        :modelValue="editingUser"
        @update:modelValue="val => editingUser.value = val"
        :loading="formLoading"
        @submit="submitEdit"
        @cancel="closeEditModal"
      />
    </el-dialog>
  </div>
</template>
