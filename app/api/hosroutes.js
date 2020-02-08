const express = require('express')
const network = require('./contract')

const routes = express.Router()
routes.use(express.json())
routes.post('/admin/getdob',async (req,res)=>{
    try {
        rbody = req.body
        const contract = await network.contract()
        const response = await contract.submitTransaction("get_dob_cert_fileName",rbody.name,rbody.p_address,rbody.c_address,rbody.phone,rbody.parent,rbody.doctor,rbody.dob)
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