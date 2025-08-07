/*
|--------------------------------------------------------------------------
| Routes file
|--------------------------------------------------------------------------
|
| The routes file is used for defining the HTTP routes.
|
*/

import { HttpContext } from '@adonisjs/core/http'
import router from '@adonisjs/core/services/router'

router.get('/', async () => {
  return {
    result: 'Ok',
  }
})
router.get('/db/', async ({ response }: HttpContext) => {
  return response.notImplemented({
    result: 'Not implemented!',
  })
})
router.get('/chaos/', async ({ response }: HttpContext) => {
  return response.notImplemented({
    result: 'Not implemented!',
  })
})
router.get('/health/', async () => {
  return {
    result: 'Ok',
  }
})
