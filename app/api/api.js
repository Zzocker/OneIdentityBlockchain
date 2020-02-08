const express = require('express')
const routes = require('./uidai')
const hosroutes = require('./hosroutes')

const PORT = "8888"
const api = express()
const logger= (req,res,next)=>{
    console.log(`${req.protocol}://${req.get('host')}${req.originalUrl}`)
    next()
}
api.use(express.json())

api.use(logger)
api.use('/uidai',routes)
api.use('/hospital',hosroutes)
api.listen(PORT,()=>{
    console.log(`Listening on port : ${PORT}`)
})
