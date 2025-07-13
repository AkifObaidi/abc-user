<script setup>
import { ref, watch, defineProps, defineEmits } from 'vue';
import { ElMessage } from 'element-plus';
import { useForm, useField } from 'vee-validate';
import * as yup from 'yup';

const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({}),
  },
  loading: {
    type: Boolean,
    default: false,
  },
});

const emit = defineEmits(['submit', 'cancel']);

const validationSchema = yup.object({
  name: yup
    .string()
    .required('Name is required')
    .min(3, 'Name must be at least 3 characters')
    .max(50, 'Name cannot exceed 50 characters'),
  email: yup
    .string()
    .required('Email is required')
    .email('Please enter a valid email address'),
  age: yup
    .number()
    .required('Age is required')
    .typeError('Age must be a number')
    .min(18, 'Age must be 18 or greater'),
});

const { handleSubmit, resetForm, setValues, errors } = useForm({
  validationSchema,
  initialValues: props.modelValue,
});

const { value: name, errorMessage: nameError } = useField('name');
const { value: email, errorMessage: emailError } = useField('email');
const { value: age, errorMessage: ageError } = useField('age');

watch(() => props.modelValue, (newVal) => {
  setValues(newVal);
}, { deep: true, immediate: true });

const submitForm = handleSubmit(async (values) => {
  emit('submit', values);
}, (context) => {
  console.error('VeeValidate form validation failed:', context.errors);
  ElMessage.error('Please correct the form errors.');
});

const cancelForm = () => {
  emit('cancel');
  resetForm();
};
</script>

<template>
  <el-form label-width="80px" @submit.prevent="submitForm">
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
      <el-button type="primary" @click="submitForm" :loading="loading">Save</el-button>
      <el-button @click="cancelForm">Cancel</el-button>
    </el-form-item>
  </el-form>
</template>