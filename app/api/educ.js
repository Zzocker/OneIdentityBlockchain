const express = require('express')
const network = require('./contract')

const routes = express.Router()
routes.use(express.json())
routes.post('/admin/addquali',async (req,res)=>{
    try {
        rbody = req.body
        const contract = await network.contract()
        const response = await contract.submitTransaction("addQualification",rbody.i_key,rbody.quali_type)
        res.status(200).json({
            result : JSON.parse(response)
        })
    } catch (error) {
        res.status(500).json({
            msg: error
        }) 
    }
})
routes.get('/getquali',async (req,res)=>{
    try {
        rbody = req.body
        const contract = await network.contract()
        const response = await contract.evaluateTransaction("getEduc",rbody.i_key)
        res.status(200).json({
            result : JSON.parse(response)
        })
    } catch (error) {
        res.status(500).json({
            msg: error
        }) 
    }
})
routes.get('/getThequali',async (req,res)=>{
    try {
        rbody = req.body
        const contract = await network.contract()
        const response = await contract.evaluateTransaction("getStateByte",rbody.q_key)
        res.status(200).json({
            result : JSON.parse(response)
        })
    } catch (error) {
        res.status(500).json({
            msg: error
        }) 
    }
})
module.exports = routes