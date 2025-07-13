<script setup>
import './NewUserView.css'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { useForm, useField } from 'vee-validate'
import * as yup from 'yup'
import userApi from '@/api/users'

const router = useRouter()

const formLoading = ref(false)

const validationSchema = yup.object({
  name: yup
    .string()
    .required('Name is required')
    .min(3, 'Name must be at least 3 characters')
    .max(50, 'Name cannot exceed 50 characters'),
  email: yup.string().required('Email is required').email('Please enter a valid email address'),
  age: yup
    .number()
    .required('Age is required')
    .typeError('Age must be a number')
    .min(18, 'Age must be 18 or greater'),
})

const { handleSubmit, resetForm } = useForm({
  validationSchema,
  initialValues: {
    name: '',
    email: '',
    age: null,
  },
})

const { value: name, errorMessage: nameError } = useField('name')
const { value: email, errorMessage: emailError } = useField('email')
const { value: age, errorMessage: ageError } = useField('age')

const createUser = handleSubmit(
  async (values) => {
    formLoading.value = true
    try {
      await userApi.createUser(values)
      ElMessage.success('User created successfully!')
      resetForm()
      router.push('/')
    } catch (error) {
      ElMessage.error('Failed to create user.')
      console.error('Error creating user:', error)
    } finally {
      formLoading.value = false
    }
  },
  (context) => {
    console.error('Form validation failed:', context.errors)
    ElMessage.error('Please correct the form errors.')
  },
)

const goBack = () => {
  router.push('/')
}
</script>

<template>
  <div class="new-user-page">
    <div class="new-user-container">
      <div class="header">
        <h1>Create New User</h1>
        <el-button @click="goBack">Back to Home</el-button>
      </div>

      <el-card class="form-card" shadow="hover">
        <el-form label-width="100px" @submit.prevent="createUser">
          <el-form-item label="Name" :error="nameError">
            <el-input v-model="name"></el-input>
          </el-form-item>

          <el-form-item label="Email" :error="emailError">
            <el-input v-model="email"></el-input>
          </el-form-item>

          <el-form-item label="Age" :error="ageError">
            <el-input-number v-model="age" :min="0" :max="120"></el-input-number>
          </el-form-item>

          <el-form-item>
            <el-button type="primary" @click="createUser" :loading="formLoading"
              >Create User</el-button
            >
            <el-button @click="resetForm">Clear Form</el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </div>
  </div>
</template>
