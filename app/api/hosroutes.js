const express = require('express')
const network = require('./contract')

const routes = express.Router()
routes.use(express.json())
routes.post('/admin/getdob',async (req,res)=>{
    try {
        rbody = req.body
        const contract = await network.contract()
        console.log(rbody)
        const response = await contract.submitTransaction("get_dob_cert_fileName",rbody.name,rbody.p_address,rbody.c_address,rbody.phone,rbody.parent,rbody.doctor,rbody.dob)
        res.status(200).json({
            status: 200,
            result : JSON.parse(response)
        })
    } catch (error) {
        res.status(500).json({
            status:500,
            msg: error
        }) 
    }
})
routes.get('/user/me/Health',async (req,res)=>{
    try {
        const contract = await network.contract()
        const response = await contract.evaluateTransaction("getHealthc",req.body.i_key)
        res.status(200).json({  
            status: 200, 
            result: JSON.parse(response)
        })
    } catch (error) {
        res.status(500).json({
            status: 500,
            msg: error.message
        })
    }
})
routes.post('/admin/addreports',async (req,res)=>{
    try {
        const contract = await network.contract()
        const response = await contract.submitTransaction("addHealthReports",req.body.i_key,req.body.doctor,req.body.type)
        res.status(200).json({   
            status: 200,
            result: JSON.parse(response)
        })
    } catch (error) {
        res.status(500).json({
            status: 500,
            msg: error.message
        })
    }
})
module.exports = routes