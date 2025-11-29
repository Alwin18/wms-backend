# Struktur Service dan API Endpoints - Warehouse & Logistics Management System

Struktur service dan endpoint lengkap untuk WMS menggunakan Golang:

# ARSITEKTUR SERVICE

```
wms-backend/
    ├── cmd/
    │   └── main.go
    ├── config/
    │   ├── app.go
    │   ├── config.go
    │   ├── fiber.go
    │   ├── gorm.go
    │   ├── logrus.go
    │   └── validator.go
    ├── internal/
    │   ├── dto
    │   ├── handlers
    │   ├── models
    │   ├── routes
    │   ├── services/
    │   │   ├── auth
    │   │   ├── master-data
    │   │   ├── inventory
    │   │   ├── inbound
    │   │   ├── outbound
    │   │   ├── movement
    │   │   ├── stock-opname
    │   │   ├── returns
    │   │   ├── production
    │   │   ├── reporting
    │   │   └── notification
    │   └── utils
    └── pkg/
        └── middleware/
            └── middleware.go
```

## 1. AUTH SERVICE

Mengelola autentikasi, autorisasi, dan manajemen user.
Endpoints:

- Authentication
  ```
  POST   /api/v1/auth/register
  POST   /api/v1/auth/login
  POST   /api/v1/auth/logout
  POST   /api/v1/auth/refresh-token
  POST   /api/v1/auth/forgot-password
  POST   /api/v1/auth/reset-password
  POST   /api/v1/auth/verify-email
  ```
- User Management
  ```
  GET    /api/v1/users
  GET    /api/v1/users/:id
  POST   /api/v1/users
  PUT    /api/v1/users/:id
  DELETE /api/v1/users/:id
  PATCH  /api/v1/users/:id/status
  PATCH  /api/v1/users/:id/password
  GET    /api/v1/users/:id/warehouses
  ```
- Role & Permission Management
  ```
  GET    /api/v1/roles
  GET    /api/v1/roles/:id
  POST   /api/v1/roles
  PUT    /api/v1/roles/:id
  DELETE /api/v1/roles/:id
  GET    /api/v1/permissions
  POST   /api/v1/roles/:id/permissions
  GET    /api/v1/roles/:id/permissions
  ```
- Audit Log
  ```
  GET    /api/v1/audit-logs
  GET    /api/v1/audit-logs/:id
  GET    /api/v1/audit-logs/user/:userId
  GET    /api/v1/audit-logs/export
  ```

## 2. MASTER DATA SERVICE

Mengelola semua data master sistem.

Endpoints:

- Product Management
  ```
  GET    /api/v1/products
  GET    /api/v1/products/:id
  POST   /api/v1/products
  PUT    /api/v1/products/:id
  DELETE /api/v1/products/:id
  PATCH  /api/v1/products/:id/status
  GET    /api/v1/products/search?q={query}
  GET    /api/v1/products/barcode/:barcode
  POST   /api/v1/products/bulk-import
  GET    /api/v1/products/export
  POST   /api/v1/products/:id/images
  DELETE /api/v1/products/:id/images/:imageId
  ```
- Product Category
  ```
  GET    /api/v1/categories
  GET    /api/v1/categories/:id
  POST   /api/v1/categories
  PUT    /api/v1/categories/:id
  DELETE /api/v1/categories/:id
  GET    /api/v1/categories/tree
  ```
- Unit of Measure (UOM)
  ```
  GET    /api/v1/uoms
  GET    /api/v1/uoms/:id
  POST   /api/v1/uoms
  PUT    /api/v1/uoms/:id
  DELETE /api/v1/uoms/:id
  GET    /api/v1/products/:productId/uoms
  POST   /api/v1/products/:productId/uoms
  ```
- Warehouse Management
  ```
  GET    /api/v1/warehouses
  GET    /api/v1/warehouses/:id
  POST   /api/v1/warehouses
  PUT    /api/v1/warehouses/:id
  DELETE /api/v1/warehouses/:id
  PATCH  /api/v1/warehouses/:id/status
  GET    /api/v1/warehouses/:id/capacity
  ```
- Location/Bin Management
  ```
  GET    /api/v1/warehouses/:warehouseId/locations
  GET    /api/v1/locations/:id
  POST   /api/v1/warehouses/:warehouseId/locations
  PUT    /api/v1/locations/:id
  DELETE /api/v1/locations/:id
  GET    /api/v1/locations/search
  GET    /api/v1/locations/:id/capacity
  PATCH  /api/v1/locations/:id/status
  POST   /api/v1/locations/bulk-create
  ```
- Zone Management
  ```
  GET    /api/v1/warehouses/:warehouseId/zones
  GET    /api/v1/zones/:id
  POST   /api/v1/warehouses/:warehouseId/zones
  PUT    /api/v1/zones/:id
  DELETE /api/v1/zones/:id
  ```
- Supplier Management
  ```
  GET    /api/v1/suppliers
  GET    /api/v1/suppliers/:id
  POST   /api/v1/suppliers
  PUT    /api/v1/suppliers/:id
  DELETE /api/v1/suppliers/:id
  GET    /api/v1/suppliers/search?q={query}
  ```
- Customer Management
  ```
  GET    /api/v1/customers
  GET    /api/v1/customers/:id
  POST   /api/v1/customers
  PUT    /api/v1/customers/:id
  DELETE /api/v1/customers/:id
  GET    /api/v1/customers/search?q={query}
  ```
- Courier/Expedition Management
  ```
  GET    /api/v1/couriers
  GET    /api/v1/couriers/:id
  POST   /api/v1/couriers
  PUT    /api/v1/couriers/:id
  DELETE /api/v1/couriers/:id
  GET    /api/v1/couriers/:id/services
  POST   /api/v1/couriers/:id/services
  ```

## 3. INVENTORY SERVICE

Mengelola stok dan pergerakan inventory.
Endpoints:

- Stock Management

```
GET    /api/v1/inventory/stock
GET    /api/v1/inventory/stock/summary
GET    /api/v1/inventory/stock/product/:productId
GET    /api/v1/inventory/stock/warehouse/:warehouseId
GET    /api/v1/inventory/stock/location/:locationId
GET    /api/v1/inventory/stock/batch/:batchNumber
GET    /api/v1/inventory/stock/serial/:serialNumber
POST   /api/v1/inventory/stock/check-availability
```

- Stock Card

```
GET    /api/v1/inventory/stock-card/product/:productId
GET    /api/v1/inventory/stock-card/warehouse/:warehouseId/product/:productId
GET    /api/v1/inventory/stock-card/export
```

- Stock Adjustment

```
GET    /api/v1/inventory/adjustments
GET    /api/v1/inventory/adjustments/:id
POST   /api/v1/inventory/adjustments
PUT    /api/v1/inventory/adjustments/:id
DELETE /api/v1/inventory/adjustments/:id
PATCH  /api/v1/inventory/adjustments/:id/approve
PATCH  /api/v1/inventory/adjustments/:id/reject
```

- Stock Reservation

```
GET    /api/v1/inventory/reservations
GET    /api/v1/inventory/reservations/:id
POST   /api/v1/inventory/reservations
DELETE /api/v1/inventory/reservations/:id
PATCH  /api/v1/inventory/reservations/:id/release
GET    /api/v1/inventory/reservations/product/:productId
```

- Stock Valuation

```
GET    /api/v1/inventory/valuation
GET    /api/v1/inventory/valuation/warehouse/:warehouseId
GET    /api/v1/inventory/valuation/product/:productId
GET    /api/v1/inventory/valuation/export
```

- Batch/Lot Management

```
GET    /api/v1/inventory/batches
GET    /api/v1/inventory/batches/:id
POST   /api/v1/inventory/batches
GET    /api/v1/inventory/batches/expiring?days={days}
GET    /api/v1/inventory/batches/expired
```

- Serial Number Management

```
GET    /api/v1/inventory/serials
GET    /api/v1/inventory/serials/:serialNumber
POST   /api/v1/inventory/serials
GET    /api/v1/inventory/serials/product/:productId
```

- Stock Alert

```
GET    /api/v1/inventory/alerts/minimum-stock
GET    /api/v1/inventory/alerts/maximum-stock
GET    /api/v1/inventory/alerts/expiry
POST   /api/v1/inventory/alerts/configure
```

## 4. INBOUND SERVICE

Mengelola penerimaan barang masuk.
Endpoints:

- Purchase Order (Reference)

```
GET    /api/v1/inbound/purchase-orders
GET    /api/v1/inbound/purchase-orders/:id
POST   /api/v1/inbound/purchase-orders
GET    /api/v1/inbound/purchase-orders/pending
```

- Advanced Shipping Notice (ASN)

```
GET    /api/v1/inbound/asn
GET    /api/v1/inbound/asn/:id
POST   /api/v1/inbound/asn
PUT    /api/v1/inbound/asn/:id
DELETE /api/v1/inbound/asn/:id
```

- Goods Receipt

```
GET    /api/v1/inbound/receipts
GET    /api/v1/inbound/receipts/:id
POST   /api/v1/inbound/receipts
PUT    /api/v1/inbound/receipts/:id
DELETE /api/v1/inbound/receipts/:id
PATCH  /api/v1/inbound/receipts/:id/complete
POST   /api/v1/inbound/receipts/:id/items
PUT    /api/v1/inbound/receipts/:id/items/:itemId
DELETE /api/v1/inbound/receipts/:id/items/:itemId
GET    /api/v1/inbound/receipts/:id/print
```

- Quality Control

```
GET    /api/v1/inbound/qc
GET    /api/v1/inbound/qc/:id
POST   /api/v1/inbound/qc
PATCH  /api/v1/inbound/qc/:id/approve
PATCH  /api/v1/inbound/qc/:id/reject
POST   /api/v1/inbound/qc/:id/photos
```

- Putaway Task

```
GET    /api/v1/inbound/putaway/tasks
GET    /api/v1/inbound/putaway/tasks/:id
POST   /api/v1/inbound/putaway/tasks
PATCH  /api/v1/inbound/putaway/tasks/:id/assign
PATCH  /api/v1/inbound/putaway/tasks/:id/start
PATCH  /api/v1/inbound/putaway/tasks/:id/complete
GET    /api/v1/inbound/putaway/tasks/user/:userId
GET    /api/v1/inbound/putaway/recommendations
```

## 5. OUTBOUND SERVICE

Mengelola proses pengiriman barang keluar.
Endpoints:

- Sales Order Management

```
GET    /api/v1/outbound/sales-orders
GET    /api/v1/outbound/sales-orders/:id
POST   /api/v1/outbound/sales-orders
PUT    /api/v1/outbound/sales-orders/:id
DELETE /api/v1/outbound/sales-orders/:id
PATCH  /api/v1/outbound/sales-orders/:id/cancel
GET    /api/v1/outbound/sales-orders/pending
```

- Stock Allocation

```
POST   /api/v1/outbound/allocations
GET    /api/v1/outbound/allocations/:orderId
DELETE /api/v1/outbound/allocations/:orderId
POST   /api/v1/outbound/allocations/batch
```

- Wave Planning

```
GET    /api/v1/outbound/waves
GET    /api/v1/outbound/waves/:id
POST   /api/v1/outbound/waves
PUT    /api/v1/outbound/waves/:id
DELETE /api/v1/outbound/waves/:id
PATCH  /api/v1/outbound/waves/:id/release
POST   /api/v1/outbound/waves/:id/orders
DELETE /api/v1/outbound/waves/:id/orders/:orderId
```

- Picking Management

```
GET    /api/v1/outbound/picking/tasks
GET    /api/v1/outbound/picking/tasks/:id
POST   /api/v1/outbound/picking/tasks
PATCH  /api/v1/outbound/picking/tasks/:id/assign
PATCH  /api/v1/outbound/picking/tasks/:id/start
PATCH  /api/v1/outbound/picking/tasks/:id/complete
POST   /api/v1/outbound/picking/tasks/:id/items/:itemId/pick
GET    /api/v1/outbound/picking/tasks/user/:userId
GET    /api/v1/outbound/picking/batch-tasks
POST   /api/v1/outbound/picking/batch-tasks
```

- Packing Management

```
GET    /api/v1/outbound/packing/stations
GET    /api/v1/outbound/packing/tasks
GET    /api/v1/outbound/packing/tasks/:id
POST   /api/v1/outbound/packing/tasks/:id/start
POST   /api/v1/outbound/packing/tasks/:id/verify-item
POST   /api/v1/outbound/packing/tasks/:id/complete
POST   /api/v1/outbound/packing/tasks/:id/packages
GET    /api/v1/outbound/packing/tasks/:id/label
```

- Shipping Management

```
GET    /api/v1/outbound/shipments
GET    /api/v1/outbound/shipments/:id
POST   /api/v1/outbound/shipments
PATCH  /api/v1/outbound/shipments/:id/dispatch
POST   /api/v1/outbound/shipments/:id/generate-awb
GET    /api/v1/outbound/shipments/:id/label
GET    /api/v1/outbound/shipments/:id/manifest
GET    /api/v1/outbound/shipments/courier/:courierId
```

- Delivery Order

```
GET    /api/v1/outbound/delivery-orders
GET    /api/v1/outbound/delivery-orders/:id
POST   /api/v1/outbound/delivery-orders
GET    /api/v1/outbound/delivery-orders/:id/print
PATCH  /api/v1/outbound/delivery-orders/:id/confirm-delivery
POST   /api/v1/outbound/delivery-orders/:id/pod
```

## 6. MOVEMENT SERVICE

Mengelola perpindahan internal barang.
Endpoints:

- Internal Transfer

```
GET    /api/v1/movements/transfers
GET    /api/v1/movements/transfers/:id
POST   /api/v1/movements/transfers
PUT    /api/v1/movements/transfers/:id
DELETE /api/v1/movements/transfers/:id
PATCH  /api/v1/movements/transfers/:id/approve
PATCH  /api/v1/movements/transfers/:id/execute
PATCH  /api/v1/movements/transfers/:id/complete
```

- Inter-Warehouse Transfer

```
GET    /api/v1/movements/inter-warehouse
GET    /api/v1/movements/inter-warehouse/:id
POST   /api/v1/movements/inter-warehouse
PATCH  /api/v1/movements/inter-warehouse/:id/dispatch
PATCH  /api/v1/movements/inter-warehouse/:id/receive
GET    /api/v1/movements/inter-warehouse/:id/print
```

- Replenishment

```
GET    /api/v1/movements/replenishments
GET    /api/v1/movements/replenishments/:id
POST   /api/v1/movements/replenishments
PATCH  /api/v1/movements/replenishments/:id/execute
GET    /api/v1/movements/replenishments/suggestions
```

- Repack/Conversion

```
GET    /api/v1/movements/repacks
GET    /api/v1/movements/repacks/:id
POST   /api/v1/movements/repacks
PATCH  /api/v1/movements/repacks/:id/execute
```

## 7. STOCK OPNAME SERVICE

Mengelola stock taking dan cycle count.
Endpoints:

- Stock Opname Planning

```
GET    /api/v1/stock-opname/plans
GET    /api/v1/stock-opname/plans/:id
POST   /api/v1/stock-opname/plans
PUT    /api/v1/stock-opname/plans/:id
DELETE /api/v1/stock-opname/plans/:id
PATCH  /api/v1/stock-opname/plans/:id/activate
```

- Cycle Count

```
GET    /api/v1/stock-opname/cycle-counts
GET    /api/v1/stock-opname/cycle-counts/:id
POST   /api/v1/stock-opname/cycle-counts
GET    /api/v1/stock-opname/cycle-counts/scheduled
```

- Count Execution

```
GET    /api/v1/stock-opname/counts
GET    /api/v1/stock-opname/counts/:id
POST   /api/v1/stock-opname/counts/:id/items
PUT    /api/v1/stock-opname/counts/:id/items/:itemId
PATCH  /api/v1/stock-opname/counts/:id/submit
GET    /api/v1/stock-opname/counts/assigned/:userId
```

- Variance & Adjustment

```
GET    /api/v1/stock-opname/variances
GET    /api/v1/stock-opname/variances/:id
PATCH  /api/v1/stock-opname/variances/:id/approve
PATCH  /api/v1/stock-opname/variances/:id/reject
POST   /api/v1/stock-opname/variances/:id/adjust
GET    /api/v1/stock-opname/variances/export
```

## 8. RETURNS SERVICE

Mengelola return dari customer dan ke supplier.
Endpoints:

- Sales Return (from Customer)

```
GET    /api/v1/returns/sales
GET    /api/v1/returns/sales/:id
POST   /api/v1/returns/sales
PUT    /api/v1/returns/sales/:id
PATCH  /api/v1/returns/sales/:id/receive
PATCH  /api/v1/returns/sales/:id/qc
PATCH  /api/v1/returns/sales/:id/complete
POST   /api/v1/returns/sales/:id/items
```

- RMA (Return Merchandise Authorization)

```
GET    /api/v1/returns/rma
GET    /api/v1/returns/rma/:id
POST   /api/v1/returns/rma
PATCH  /api/v1/returns/rma/:id/approve
PATCH  /api/v1/returns/rma/:id/reject
GET    /api/v1/returns/rma/:id/print
```

- Purchase Return (to Supplier)

```
GET    /api/v1/returns/purchase
GET    /api/v1/returns/purchase/:id
POST   /api/v1/returns/purchase
PATCH  /api/v1/returns/purchase/:id/dispatch
GET    /api/v1/returns/purchase/:id/print
```

- Return Disposition

```
GET    /api/v1/returns/dispositions
POST   /api/v1/returns/:returnId/disposition
PATCH  /api/v1/returns/dispositions/:id/execute
```

## 9. PRODUCTION SERVICE

Mengelola assembly/disassembly sederhana.
Endpoints:

- Bill of Materials (BOM)

```
GET    /api/v1/production/bom
GET    /api/v1/production/bom/:id
POST   /api/v1/production/bom
PUT    /api/v1/production/bom/:id
DELETE /api/v1/production/bom/:id
GET    /api/v1/production/bom/product/:productId
POST   /api/v1/production/bom/:id/components
DELETE /api/v1/production/bom/:id/components/:componentId
```

- Work Order

```
GET    /api/v1/production/work-orders
GET    /api/v1/production/work-orders/:id
POST   /api/v1/production/work-orders
PUT    /api/v1/production/work-orders/:id
DELETE /api/v1/production/work-orders/:id
PATCH  /api/v1/production/work-orders/:id/start
PATCH  /api/v1/production/work-orders/:id/complete
PATCH  /api/v1/production/work-orders/:id/cancel
```

- Material Issue

```
POST   /api/v1/production/work-orders/:id/issue-materials
GET    /api/v1/production/work-orders/:id/issued-materials
Production Receipt
POST   /api/v1/production/work-orders/:id/receive-finished-goods
GET    /api/v1/production/work-orders/:id/receipts
```

## 10. INTEGRATION SERVICE

Mengelola integrasi dengan sistem eksternal.
Endpoints:

- ERP Integration

```
POST   /api/v1/integrations/erp/sync-products
POST   /api/v1/integrations/erp/sync-customers
POST   /api/v1/integrations/erp/sync-suppliers
POST   /api/v1/integrations/erp/sync-purchase-orders
POST   /api/v1/integrations/erp/sync-sales-orders
POST   /api/v1/integrations/erp/push-stock-movements
POST   /api/v1/integrations/erp/push-inventory-valuation
GET    /api/v1/integrations/erp/sync-status
```

- Marketplace Integration

```
POST   /api/v1/integrations/marketplace/sync-products
POST   /api/v1/integrations/marketplace/sync-orders
POST   /api/v1/integrations/marketplace/update-stock
POST   /api/v1/integrations/marketplace/update-order-status
GET    /api/v1/integrations/marketplace/channels
POST   /api/v1/integrations/marketplace/channels
```

- Courier Integration

```
POST   /api/v1/integrations/courier/check-rates
POST   /api/v1/integrations/courier/create-booking
POST   /api/v1/integrations/courier/track-shipment
POST   /api/v1/integrations/courier/cancel-booking
GET    /api/v1/integrations/courier/tracking/:awb
```

- Webhook Management

```
GET    /api/v1/integrations/webhooks
GET    /api/v1/integrations/webhooks/:id
POST   /api/v1/integrations/webhooks
PUT    /api/v1/integrations/webhooks/:id
DELETE /api/v1/integrations/webhooks/:id
GET    /api/v1/integrations/webhooks/:id/logs
POST   /api/v1/integrations/webhooks/:id/test
```

- API Keys

```
GET    /api/v1/integrations/api-keys
POST   /api/v1/integrations/api-keys
DELETE /api/v1/integrations/api-keys/:id
PATCH  /api/v1/integrations/api-keys/:id/regenerate
```

## 11. REPORTING SERVICE

Mengelola laporan dan analytics.
Endpoints:

- Inventory Reports

```
GET    /api/v1/reports/inventory/stock-summary
GET    /api/v1/reports/inventory/stock-by-warehouse
GET    /api/v1/reports/inventory/stock-by-location
GET    /api/v1/reports/inventory/stock-valuation
GET    /api/v1/reports/inventory/stock-aging
GET    /api/v1/reports/inventory/stock-movement
GET    /api/v1/reports/inventory/slow-moving
GET    /api/v1/reports/inventory/fast-moving
GET    /api/v1/reports/inventory/dead-stock
GET    /api/v1/reports/inventory/expiring-products
```

- Warehouse Performance

```
GET    /api/v1/reports/warehouse/order-fulfillment
GET    /api/v1/reports/warehouse/picking-performance
GET    /api/v1/reports/warehouse/packing-performance
GET    /api/v1/reports/warehouse/receiving-performance
GET    /api/v1/reports/warehouse/productivity
GET    /api/v1/reports/warehouse/shrinkage
GET    /api/v1/reports/warehouse/space-utilization
GET    /api/v1/reports/warehouse/accuracy-rate
```

- Operational Reports

```
GET    /api/v1/reports/operations/inbound-summary
GET    /api/v1/reports/operations/outbound-summary
GET    /api/v1/reports/operations/returns-summary
GET    /api/v1/reports/operations/transfers-summary
GET    /api/v1/reports/operations/stock-adjustments
```

- Analytics Dashboard

```
GET    /api/v1/reports/dashboard/overview
GET    /api/v1/reports/dashboard/kpi
GET    /api/v1/reports/dashboard/alerts
GET    /api/v1/reports/dashboard/trends
```

- Custom Reports

```
GET    /api/v1/reports/custom
POST   /api/v1/reports/custom
GET    /api/v1/reports/custom/:id/execute
DELETE /api/v1/reports/custom/:id
```

- Export

```
POST   /api/v1/reports/export/excel
POST   /api/v1/reports/export/pdf
POST   /api/v1/reports/export/csv
GET    /api/v1/reports/export/status/:jobId
GET    /api/v1/reports/export/download/:jobId
```

## 12. NOTIFICATION SERVICE

Mengelola notifikasi dan alert.
Endpoints:

- Notification Management

```
GET    /api/v1/notifications
GET    /api/v1/notifications/:id
PATCH  /api/v1/notifications/:id/read
PATCH  /api/v1/notifications/read-all
DELETE /api/v1/notifications/:id
GET    /api/v1/notifications/unread-count
```

- Alert Configuration

```
GET    /api/v1/notifications/alerts/config
POST   /api/v1/notifications/alerts/config
PUT    /api/v1/notifications/alerts/config/:id
DELETE /api/v1/notifications/alerts/config/:id
```

- Email Templates

```
GET    /api/v1/notifications/email-templates
GET    /api/v1/notifications/email-templates/:id
POST   /api/v1/notifications/email-templates
PUT    /api/v1/notifications/email-templates/:id
```

## Standart Success Response

```
{
  "success": true,
  "message": "Success message",
  "data": {},
  "meta": {
    "page": 1,
    "limit": 10,
    "total": 100,
    "totalPages": 10
  },
  "errors": []
}
```

## Standart Error Response

```
{
  "success": false,
  "message": "Error message",
  "errors": [
    "message": "Email already exists"
  ]
}
```

## Authentication

- JWT Bearer Token
- Header: Authorization: Bearer `<token>`
- Refresh token mechanism untuk session management

## Versioning

- API versioning menggunakan URL path (/api/v1/)
- Backward compatibility untuk versi lama

---

# Deskripsi Detail Service - Warehouse & Logistics Management System

## 1. AUTH SERVICE

- Tujuan:
  Menyediakan sistem autentikasi dan autorisasi yang aman, mengelola identitas pengguna, dan mengontrol akses ke berbagai fitur sistem berdasarkan role dan permission.
- Proses Bisnis:
  1. User Registration & Onboarding
     - Admin/HR menambahkan user baru dengan data lengkap (nama, email, jabatan, gudang yang di-assign)
     - Sistem mengirim email verifikasi
     - User melakukan verifikasi dan set password pertama kali
     - User otomatis di-assign ke default role sesuai jabatan
  2. Authentication Flow
     - User login dengan email/username dan password
     - Sistem validasi kredensial dan generate JWT token (access + refresh token)
     - Access token berlaku singkat (15-30 menit), refresh token lebih lama (7-30 hari)
     - Setiap request ke API menggunakan access token di header
     - Ketika access token expired, client menggunakan refresh token untuk dapat access token baru
  3. Authorization & Access Control
     - Setiap endpoint dilindungi middleware yang cek token validity
     - Sistem cek role dan permission user untuk akses fitur tertentu
     - Permission berbasis resource + action (contoh: "inventory.read", "sales_order.create")
     - User bisa punya multiple roles atau custom permission
     - Akses bisa dibatasi per warehouse (user hanya bisa akses data gudang tertentu)
  4. Role Management
     - Predefined roles: System Admin, Warehouse Manager, Supervisor, Picker, Packer, Inventory Planner, Finance, Customer Service
     - Setiap role punya set permission default
     - Admin bisa create custom role dan assign permission sesuai kebutuhan
     - Role bisa di-assign ke multiple users
  5. Audit & Security
     - Semua aktivitas login/logout tercatat dengan timestamp, IP address, device info
     - Failed login attempts di-track untuk detect brute force attack
     - Sensitive actions (delete, approval, adjustment) tercatat di audit log
     - Password policy enforcement (minimum length, complexity, expiry)
- Goals:
  - ✅ Zero unauthorized access - hanya user terautentikasi dan terotorisasi yang bisa akses sistem
  - ✅ Audit trail lengkap untuk compliance dan forensik
  - ✅ Session management yang aman dengan token expiry otomatis
  - ✅ Granular permission control per fitur dan per gudang
  - ✅ User dapat bekerja sesuai role tanpa akses ke data/fitur yang tidak relevan
  - ✅ Deteksi dini suspicious activities (multiple failed login, unusual access pattern)

## 2. MASTER DATA SERVICE

- Tujuan:
  Menyediakan single source of truth untuk semua data master yang digunakan oleh seluruh modul WMS, memastikan konsistensi dan integritas data.
- Proses Bisnis:
  1. Product Management

     - Setup Produk Baru:
       - Admin/Inventory Planner input data produk: SKU, nama, deskripsi, kategori, brand
       - Definisi UOM (Unit of Measure) utama dan konversi (1 karton = 12 pcs)
       - Generate atau input barcode/QR code untuk scanning
       - Set dimensi (P x L x T), berat untuk perhitungan kapasitas dan ongkir
       - Tentukan apakah produk memerlukan batch/lot tracking atau serial number
       - Set tipe penyimpanan (ambient, chilled, frozen)
       - Upload foto produk untuk identifikasi visual
     - Product Categorization:
       - Organisasi produk dalam hierarki kategori (Department > Category > Sub-Category)
       - Memudahkan filtering, reporting, dan slotting strategy
       - Kategori bisa punya attribute khusus (mis: fashion punya size/color)
     - Bundle/Kit Management:
       - Produk bundle terdiri dari multiple produk lain
       - System track stok komponen dan produk bundle secara terpisah
       - Saat bundle terjual, stok komponen berkurang otomatis
  2. Warehouse & Location Management

     - Warehouse Setup:
       - Pendefinisian gudang baru dengan alamat lengkap, tipe (DC/warehouse/store)
       - Set timezone, jam operasional, contact person
       - Konfigurasi default rules (FIFO/FEFO, wave picking strategy)
     - Location Hierarchy:
       - Struktur: Warehouse → Zone → Aisle → Rack → Level → Bin
       - Setiap lokasi punya kode unik untuk scanning (contoh: WH1-A-01-A-01)
       - Generate barcode untuk setiap lokasi
     - Location Attributes:
       - Tipe lokasi: Receiving (temp), Storage (bulk), Picking (forward), Packing, Damaged, Return Hold, Dispatch
       - Kapasitas: max qty, max volume (m³), max weight (kg)
       - Restriction: produk tertentu tidak boleh dicampur
       - Temperature zone untuk produk cold chain
     - Slotting Strategy (Fase 2):
       - Fast-moving items ditempatkan dekat area picking
       - Slow-moving di area storage lebih jauh
       - Produk heavy/bulky di level bawah
       - Sistem beri rekomendasi lokasi optimal saat putaway
  3. Supplier & Customer Management

     - Data lengkap kontak, alamat, terms pembayaran
     - Rating supplier (delivery performance, quality)
     - Credit limit untuk customer
     - Integration dengan sistem akuntansi untuk invoice reconciliation
  4. Courier Management

     - Daftar jasa ekspedisi yang digunakan (3PL atau internal fleet)
     - Service level per courier (Regular, Express, Same Day)
     - Konfigurasi API credentials untuk auto booking dan tracking
     - Rate card atau integration dengan shipping aggregator untuk perhitungan ongkir
- Goals:
  - ✅ Single source of truth - tidak ada duplikasi atau inkonsistensi data master
  - ✅ Data master akurat dan up-to-date untuk operasional harian
  - ✅ Kemudahan pencarian produk via SKU, barcode, atau nama
  - ✅ Struktur lokasi yang clear untuk memudahkan operator menemukan barang
  - ✅ Fleksibilitas dalam mendefinisikan atribut produk sesuai kebutuhan bisnis
  - ✅ Mendukung expansion multi-warehouse tanpa duplikasi setup
  - ✅ Integration-ready dengan sistem eksternal (ERP, marketplace, POS)

## 3. INVENTORY SERVICE

- Tujuan:
  Memberikan visibilitas real-time terhadap posisi stok di seluruh warehouse, mengelola pergerakan inventory, dan memastikan akurasi stok.
  Proses Bisnis:

  1. Real-Time Stock Tracking
     - Multi-Dimensional Visibility:
       - Stok per produk (total across all warehouses)
       - Stok per warehouse (breakdown per lokasi/bin)
       - Stok per lokasi/bin (detail sampai level terkecil)
       - Stok per batch/lot dengan expiry date
       - Stok per serial number untuk produk high-value
     - Stock Status:
       - Available: stok yang bisa di-allocate untuk order
       - Reserved: sudah di-allocate untuk order tertentu, belum di-pick
       - Blocked: tidak bisa digunakan (pending QC, hold, investigation)
       - Damaged: rusak, tidak layak jual
       - On-Hold: temporary hold untuk keperluan tertentu
       - In-Transit: sedang dalam perjalanan antar gudang
  2. Stock Movement & Transaction
     - Setiap pergerakan stok tercatat dengan detail:
       - Transaction type (inbound, outbound, transfer, adjustment)
       - Quantity in/out
       - Location from/to
       - Reference document (PO, SO, DO, Transfer Order)
       - User yang melakukan
       - Timestamp
       - Reason/notes
     - Stock Card:
       - Riwayat lengkap setiap transaksi per produk
       - Running balance setelah setiap transaksi
       - Trace-ability dari awal sampai akhir
  3. Stock Reservation & Allocation
     - Ketika sales order masuk, sistem otomatis reserve stok:
       - Check availability per warehouse
       - Apply allocation rules (FIFO/FEFO/nearest warehouse)
       - Create reservation dengan expiry time
       - Kalau order cancel, reservation auto-release
     - Prevent overselling:
       - Available stock = on-hand - reserved - blocked
       - Sync available stock ke marketplace/e-commerce real-time
  4. Batch/Lot & Expiry Management
     - Produk dengan expiry (food, cosmetics, pharma) wajib punya batch/lot
     - Setiap receiving, input batch number dan expiry date
     - Sistem enforce FEFO (First Expired First Out):
       - Saat picking, sistem prioritaskan batch yang paling dulu expire
       - Alert untuk batch yang mendekati expiry (30/60/90 hari)
       - Block batch yang sudah expired secara otomatis
     - Recall management:
       - Kalau ada batch bermasalah, bisa di-trace semua transaksi
       - Block semua stok dari batch tersebut
       - Generate list customer yang sudah terima batch tersebut
  5. Serial Number Tracking
     - Untuk produk high-value atau regulated (elektronik, alat medis):
       - Setiap unit punya serial number unik
       - Track pergerakan per serial number dari receiving sampai shipment
       - Warranty management per serial
       - Prevent duplicate serial number entry
  6. Stock Adjustment
     - Adjustment diperlukan untuk:
       - Hasil stock opname (selisih sistem vs fisik)
       - Damage/loss
       - Quality downgrade (dari grade A ke grade B)
       - Data correction

    - Adjustment workflow:
      - Operator submit adjustment dengan reason dan supporting document
      - Supervisor/Manager review dan approve
      - Setelah approved, stok di-update dan tercatat di stock card
      - Finance di-notify untuk nilai persediaan adjustment
  7. Stock Valuation

Perhitungan nilai persediaan (Rp) menggunakan metode:

Average cost (paling umum)
FIFO costing (first-in-first-out)

Setiap pergerakan stok mempengaruhi nilai persediaan:

Inbound menambah nilai (qty × unit cost)
Outbound mengurangi nilai (qty × average/FIFO cost)
Calculate COGS (Cost of Goods Sold) per transaksi

Integration dengan accounting:

Real-time sync nilai persediaan untuk balance sheet
COGS untuk income statement
Journal entry otomatis untuk setiap transaksi

Stock Alert & Notification

Minimum Stock Alert:

Set minimum stock level per produk per warehouse
Alert ke Inventory Planner jika stok di bawah minimum
Auto-generate purchase requisition (fase lanjutan)

Maximum Stock Alert:

Alert jika stok melebihi maximum untuk avoid overstock

Expiry Alert:

Multi-level alert: 90 days, 60 days, 30 days before expiry
Notify sales team untuk flash sale atau promo
Alert warehouse untuk segregate expired stock

Goals:

✅ Stock accuracy ≥ 98% (gap antara sistem dan fisik minimal)
✅ Zero stockout untuk fast-moving items (selalu ada stok)
✅ Zero overselling (tidak menjual stok yang tidak tersedia)
✅ Minimize expired goods (rotate stock dengan FEFO)
✅ Real-time visibility untuk decision making
✅ Trace-ability lengkap untuk recall atau investigation
✅ Automated alerts untuk proactive inventory management
✅ Accurate inventory valuation untuk financial reporting

## 4. INBOUND SERVICE

Tujuan:
Mengelola penerimaan barang masuk secara efisien dan akurat, memastikan barang diterima sesuai dengan purchase order, dan menempatkan barang ke lokasi penyimpanan yang optimal.
Proses Bisnis:

Pre-Receiving (ASN - Advanced Shipping Notice)

Supplier/ERP send ASN sebelum barang tiba:

PO reference
Expected arrival date/time
List produk, quantity, batch/lot info
Packing info (jumlah pallet, carton)

Warehouse team prepare:

Allocate receiving dock
Assign staff untuk receiving
Prepare area receiving untuk barang masuk
Review PO details untuk anticipate discrepancy

Receiving Process

Truck Arrival:

Security log truck arrival time, driver info, seal number
Check surat jalan (delivery note) vs ASN/PO

Physical Inspection:

Receiver scan barcode pada setiap carton/item
Sistem auto-match dengan PO line items
Count quantity dan verifikasi kondisi fisik
Handling discrepancies:

Over-receive (terima lebih dari PO): dalam batas toleransi (misal 5%) bisa diterima, di atas itu perlu approval
Under-receive (terima kurang dari PO): mark PO line sebagai partial, sisa qty status backorder
Wrong item: reject dan buat return note ke supplier
Damage: segregate ke damaged area, photo evidence

Batch/Lot Recording:

Input batch/lot number (from supplier label atau generate internal)
Scan/input expiry date
Input manufacturing date jika diperlukan
Sistem validate expiry policy (misal: min 80% shelf life remaining)

Serial Number Recording:

Untuk produk serial, scan/input setiap serial number
Sistem validate unique serial number (no duplicate)
Associate serial dengan batch/lot

Quality Control (QC)

QC Inspection:

Sample inspection atau 100% inspection (tergantung policy)
QC checklist: physical condition, quantity accuracy, specification match
QC result: Pass, Conditional Pass (minor issues), Fail (reject)

QC Outcomes:

Pass → proceed to putaway
Conditional Pass → move to hold area, wait for approval
Fail → segregate ke rejected area, initiate return to supplier

Documentation:

Photo evidence untuk damage/defect
QC report attached ke receiving document
Notify buyer/procurement untuk corrective action

Goods Receipt Document

Sistem generate Goods Receipt Note (GRN):

GRN number (auto-increment)
PO reference
Actual received quantity per line item
Batch/lot/serial details
QC status
Receiver name dan timestamp

GRN approval workflow:

Auto-approve jika full receive dan pass QC
Require supervisor approval untuk partial atau over-receive

Integration:

Push GRN ke ERP/accounting untuk 3-way match (PO - GRN - Invoice)
Update PO status (partial/fully received)
Trigger payment process jika terms sudah match

Putaway Planning

Setelah GRN approved, sistem generate putaway tasks:

Group items berdasarkan destination zone (chilled, ambient, hazmat)
Recommend optimal location berdasarkan:

Product velocity (fast-moving → picking location, slow-moving → bulk storage)
Product characteristics (heavy items → lower level, fragile → special area)
Available capacity per location
Batch consolidation (batch yang sama ke lokasi sama jika memungkinkan)

Putaway task assignment:

Assign ke available operator via mobile app
Task include: source location (receiving), destination location, product, quantity
Operator dapat task dalam batch untuk efficiency

Putaway Execution

Mobile-Guided Putaway:

Operator login di mobile app, see assigned tasks
Navigate ke source location (receiving area)
Scan location barcode untuk confirm source
Scan product barcode untuk confirm product
Sistem confirm quantity atau operator input manual
Navigate ke destination location (sistem beri petunjuk aisle/rack)
Scan destination location untuk confirm putaway
Sistem update stok: tambah stok di destination, kurangi dari receiving area

Directed vs Undirected:

Directed: sistem kasih exact location (lebih strict)
Undirected: operator pilih location sendiri dalam zone yang sama (lebih flexible)

Exception Handling:

Location full: sistem auto-suggest alternative location
Product not match: alert operator, require re-scan atau supervisor intervention
Damage found: operator bisa report damage, move ke damaged area

Stock Update & Notification

Real-time stock update:

On-hand stock increase di destination location
Available stock increase (jika sudah pass QC dan di-putaway)
Update stock card dengan transaction detail

Notifications:

Notify Inventory Planner: new stock available
Notify Sales: backorder bisa di-fulfill sekarang
Sync available stock ke marketplace/e-commerce
Alert jika stok baru causes overstock (above maximum threshold)

Goals:

✅ Receiving accuracy 99% (qty dan quality sesuai PO)
✅ Average receiving lead time < 2 jam dari truck arrival sampai putaway complete
✅ Zero lost inventory during receiving process
✅ 100% traceability (tahu dari mana, kapan, siapa yang terima)
✅ Minimize dock congestion dengan efficient receiving process
✅ Seamless integration dengan procurement dan accounting
✅ Early detection of supplier issues (frequent short shipment, damage, wrong items)

## 5. OUTBOUND SERVICE

Tujuan:
Memproses pesanan pelanggan dengan cepat dan akurat, dari alokasi stok, picking, packing, hingga pengiriman, untuk mencapai on-time delivery dan customer satisfaction.
Proses Bisnis:

Order Entry & Validation

Sales order masuk dari berbagai channel:

ERP/sales module
Marketplace (Tokopedia, Shopee, Lazada)
E-commerce website
POS (untuk retail)
Manual entry

Order validation:

Customer info completeness
Delivery address valid
Product SKU valid (exist in master data)
Requested quantity vs available stock
Special instructions (gift wrap, COD, fragile)

Stock Allocation

Allocation Rules Engine:

Priority 1: Warehouse terdekat dengan destination (minimize shipping cost & time)
Priority 2: Warehouse dengan stock sufficient untuk full order (avoid split shipment)
Priority 3: FEFO/FIFO compliance (allocate batch yang paling dulu expire)
Priority 4: Location optimization (picking location dulu, baru storage)

Allocation Process:

Sistem cari available stock per allocation rule
Create reservation: update stock status dari available → reserved
Reservation lock selama X jam (misal 24 jam), auto-release jika order cancel
Generate pick slip dengan detail: product, quantity, source location, batch/serial if applicable

Partial Allocation:

Jika tidak cukup stok untuk full order, 2 opsi:

Split shipment: ship available qty dulu, backorder sisanya
Wait: hold order sampai full stock available

Bisnis rule menentukan mana yang dipilih

Wave Planning & Batch Picking

Wave Planning:

Group multiple orders untuk picking efficiency
Wave criteria:

Shipping cut-off time (semua order harus ship hari ini)
Destination area (semua order untuk Jakarta dalam 1 wave)
Courier/carrier (semua JNE orders dalam 1 wave)
Product zone (semua chilled products dalam 1 wave)

Picking Strategy:

Single Order Picking: 1 picker ambil 1 order complete (untuk order kecil atau urgent)
Batch Picking: 1 picker ambil multiple order sekaligus (lebih efisien untuk order dengan produk sama)
Wave Picking: multiple pickers untuk 1 wave besar, hasil di-consolidate di staging area
Zone Picking: setiap picker handle 1 zone, produk pass ke zone berikutnya (untuk warehouse besar)

Picking Execution

Picking Task Assignment:

Sistem assign picking task ke available picker
Picker receive task di mobile device
Picking list berisi: order number, product, qty, source location
List ter-sort berdasarkan optimal route (minimize walking distance)

Mobile-Guided Picking:

Picker navigate ke location pertama
Scan location barcode untuk verify
Scan product barcode untuk verify (prevent mis-pick)
System display expected qty, picker confirm atau adjust
System display batch/serial info jika applicable
Picker input actual picked qty
System confirm dan update stock real-time
Repeat untuk produk berikutnya dalam list

Pick Validation:

Computer vision (fase lanjutan): camera verify product match
Weight verification: timbang picked items vs expected weight
Short pick handling: jika stok tidak cukup, picker report discrepancy, system adjust allocation

Pick to Tote/Cart:

Untuk batch picking: setiap order punya tote/bin berbeda
Picker scan tote barcode sebelum put product
System track product in setiap tote untuk avoid mixing

Packing Station

Order Consolidation:

Picked items masuk ke packing area
Packer scan order barcode untuk retrieve order details
System display list produk yang harus ada dalam order

Packing Verification:

Packer verify setiap item:

Scan product barcode
System confirm item ada dalam order
Visual check kondisi produk
Prevent short ship atau wrong item

Packaging Selection:

System recommend box/envelope size berdasarkan dimensi produk
Packer pilih packaging material (box, bubble wrap, void fill)
Input packaging dimension dan weight (untuk ongkir calculation)

Multi-Package Handling:

Jika order tidak muat dalam 1 box, split ke multiple packages
Each package punya sub-tracking number
Customer dapat track setiap package

Packing Documentation:

Print invoice/packing slip dan masukkan ke box
Print shipping label (nama, alamat, barcode)
Attach label ke box
Seal box dan mark sebagai ready to ship

Shipping & Dispatch

Courier Integration:

Sistem connect dengan courier API (JNE, J&T, SiCepat, etc.)
Auto-generate AWB (Air Waybill / resi number)
Calculate shipping cost berdasarkan weight, dimension, destination
Print shipping label dengan AWB barcode

Manifest Creation:

Group semua shipment per courier per destination
Generate manifest document (list semua AWB untuk driver/courier)
Driver/courier scan manifest untuk acknowledge pickup

Handover Process:

Warehouse staff handover packages ke courier
Scan setiap package untuk confirm dispatch
System update order status: Shipped/In Transit
Capture signature dan timestamp

Customer Notification:

Auto-send notification ke customer:

Order shipped notification dengan AWB
Tracking link
Estimated delivery date

Notify sales team untuk follow up

Delivery Tracking & Confirmation

Real-time tracking integration dengan courier:

Update order status berdasarkan courier tracking (in transit, out for delivery, delivered)
Handle exceptions (failed delivery, return to sender)

Proof of Delivery (POD):

Capture recipient name, signature, photo
Timestamp dan GPS location
Update order status: Delivered/Completed

Post-Delivery:

Trigger invoice finalization
Update sales & revenue report
Request customer feedback/review

Goals:

✅ Order fulfillment accuracy 99.5% (correct items, correct quantity)
✅ On-time shipment rate ≥ 97% (ship sesuai SLA/cut-off time)
✅ Average pick-to-ship time < 4 jam untuk standard order
✅ Zero miss-pick atau wrong-shipment
✅ Maximize picking productivity (orders/picker/hour)
✅ Minimize shipping cost melalui optimal allocation dan packaging
✅ Real-time order visibility untuk customer dan customer service
✅ Seamless returns process untuk customer satisfaction

## 6. MOVEMENT SERVICE

Tujuan:
Mengelola perpindahan barang di dalam warehouse (internal transfer) dan antar warehouse, serta replenishment dari storage ke picking location, untuk menjaga kelancaran operasional dan optimasi layout.
Proses Bisnis:

Internal Transfer (Intra-Warehouse)

Use Cases:

Relocation: pindah produk dari lokasi tidak optimal ke lokasi lebih baik (hasil analisa velocity)
Consolidation: gabung stok dari multiple locations ke 1 location untuk space efficiency
Reorganization: warehouse layout change, butuh move banyak items
Quality segregation: move damaged goods ke damaged area
Quarantine: move suspected items ke quarantine area pending investigation

Transfer Process:

Planner/Supervisor create transfer request:

Product, quantity
Source location
Destination location
Reason/notes

Transfer request bisa butuh approval (tergantung policy)
Setelah approved, generate transfer task
Operator receive task di mobile:

Navigate ke source location
Scan location & product
Confirm qty to move
Navigate ke destination location
Scan destination location
Confirm transfer complete

System update stock real-time:

Deduct from source location
Add to destination location
Record transfer transaction in stock card

Inter-Warehouse Transfer

Use Cases:

Stock balancing: warehouse A overstock, warehouse B low stock → transfer A ke B
Demand-based allocation: ada big order dari area B, lebih cepat ship dari warehouse B
Consolidation: close warehouse, transfer all stock ke warehouse lain
Return to central DC: store/branch return unsold goods ke DC

Transfer Workflow:

Initiation:

Inventory planner analyze stock level per warehouse
Identify transfer need
Create Inter-Warehouse Transfer Order (IWTO):

Source warehouse
Destination warehouse
Product list & quantities
Expected transit time
Transportation method (internal truck, 3PL)

Approval:

Source warehouse manager approve (confirm stock available)
Destination warehouse manager approve (confirm capacity available)
Logistics manager approve transportation

Outbound from Source:

Source warehouse treat as outbound shipment
Pick, pack, generate delivery note
Load ke truck, capture seal number
Status: In Transit
Deduct stock from source warehouse (status: in-transit)

In-Transit Tracking:

Track shipment location via GPS (fase lanjutan)
Alert jika delay atau route deviation

Inbound to Destination:

Destination warehouse treat as inbound receiving
Check seal integrity
Receive & count quantity
Handle discrepancy jika ada (short/over/damage)
QC inspection
Putaway ke location
Add stock to destination warehouse
Confirm IWTO complete

Reconciliation:

System auto-reconcile quantity sent vs received
Generate variance report jika ada discrepancy
Investigation untuk loss/damage dalam transit

Replenishment (Storage to Picking)

Concept:

Bulk stock di storage location (pallet qty)
Forward stock di picking location (case/piece qty)
Picking location punya kapasitas terbatas
Perlu regular replenishment dari storage ke picking

Replenishment Triggers:

Min/Max Method:

Set minimum qty untuk picking location
Jika stock < min, trigger replenishment
Replenish sampai max level

Demand-Based:

Forecast picking demand (berdasarkan order pipeline)
Proactive replenishment sebelum stock habis

Wave-Based:

Replenish sebelum wave picking start
Ensure picking location fully stocked untuk wave efficiency

Replenishment Process:

System auto-generate replenishment task
Task include: product, qty, source (storage), destination (picking)
Assign ke operator
Operator execute: pick from storage, move to picking location
Update stock: deduct storage, add picking location

Repack/Conversion

Use Cases:

Break bulk: customer order 5 pcs, tapi stok ada dalam karton (12 pcs) → unpack karton jadi pcs
Consolidation: ada 50 loose pcs, repack jadi karton untuk save space
UOM conversion untuk inventory optimization

Repack Process:

Create repack order: source UOM, target UOM, conversion ratio
Example: 1 karton (12 pcs) → 12 pcs
System calculate:

Deduct 1 karton from stock
Add 12 pcs to stock
Record conversion transaction

Physical execution:

Operator ambil karton
Unpack dan count
Re-label dengan pcs barcode
Put ke location untuk pcs stock

Goals:

✅ Optimize space utilization (move slow-moving ke area jauh, fast-moving ke area dekat)
✅ Ensure picking location selalu stocked (zero stockout di picking location)
✅ Efficient inter-warehouse stock balancing (reduce overstock di 1 warehouse, prevent stockout di warehouse lain)
✅ Minimize handling cost untuk movement (minimize trips, optimize batch movement)
✅ 100% accuracy dalam transfer (zero lost inventory during movement)
✅ Real-time visibility stock in-transit antar warehouse
✅ Flexible UOM management untuk support berbagai order size
