<script setup>
import './UserInfoView.css'
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';
import userApi from '@/api/users';

const route = useRoute();
const router = useRouter();

const user = ref(null);
const loading = ref(true);
const error = ref(null);

const userId = route.params.id;

const loadUser = async () => {
  loading.value = true;
  error.value = null;
  try {
    const response = await userApi.getUser(userId);
    user.value = response.data;
  } catch (err) {
    error.value = 'Failed to load user information.';
    ElMessage.error('Failed to load user information.');
    console.error('Error loading user:', err);
  } finally {
    loading.value = false;
  }
};

const goBack = () => {
  router.push('/');
};

onMounted(() => {
  if (userId) {
    loadUser();
  } else {
    error.value = 'User ID not provided.';
    ElMessage.error('User ID not provided in the URL.');
    loading.value = false;
  }
});
</script>

<template>
  <div class="user-info-page">
    <div class="header">
      <h1>User Information</h1>
      <el-button @click="goBack">Back to Home</el-button>
    </div>

    <el-alert
      v-if="error"
      title="Error"
      type="error"
      :closable="false"
      style="margin-bottom: 1rem;"
    />

    <el-card v-loading="loading" class="info-card" shadow="hover">
      <template v-if="user">
        <el-descriptions
          class="margin-top"
          :column="1"
          border
        >
          <el-descriptions-item label="Name">{{ user.name }}</el-descriptions-item>
          <el-descriptions-item label="Email">{{ user.email }}</el-descriptions-item>
          <el-descriptions-item label="Age">{{ user.age }}</el-descriptions-item>
          <el-descriptions-item label="Created At">
            {{ user.createdAt ? new Date(user.createdAt).toLocaleString() : 'N/A' }}
          </el-descriptions-item>
          <el-descriptions-item label="Updated At">
            {{ user.updatedAt ? new Date(user.updatedAt).toLocaleString() : 'N/A' }}
          </el-descriptions-item>
        </el-descriptions>
      </template>
      <template v-else-if="!loading && !error">
        <el-empty description="User not found or no data available."></el-empty>
      </template>
    </el-card>
  </div>
</template>
