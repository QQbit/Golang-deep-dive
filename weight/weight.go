package weight

type KG float64
type LB float64

//package initialzation จะใช้ในกรณี package เรามีความซับซ้อน และ ต้องการ func อันหนึ่งที่ต้องการ excute มันก่อนคนอื่นจะเอาไปใช้งาน

const kgToLbRatio = 2.2046226218 /*ถ้าเกิดว่า kgToLbRation ตรงนี้มันต้องดึงข้อมูลมาจาก database หรือว่า web service
เราจะสามารถ initial package ของเราโดยใช้ฟังชั่น special อันหนึ่ง ที่ชื่อว่า initialzation */
