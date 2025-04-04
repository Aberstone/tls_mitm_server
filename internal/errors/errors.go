/*
 * Copyright (C) 2024 aberstone
 *
 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 3 of the License, or (at your option) any later version.
 *
 * This library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with this library; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA
 */
package errors

import (
	"fmt"
)

// ErrorType 定义错误类型
type ErrorType string

const (
	// ErrConfiguration 配置错误
	ErrConfiguration ErrorType = "Configuration"
	// ErrCertificate 证书错误
	ErrCertificate ErrorType = "Certificate"
	// ErrNetwork 网络错误
	ErrNetwork ErrorType = "Network"
	// ErrTLS TLS错误
	ErrTLS ErrorType = "TLS"
	// ErrProxy 代理错误
	ErrProxy ErrorType = "Proxy"
)

// ProxyError 代理服务器错误
type ProxyError struct {
	Type    ErrorType
	Message string
	Cause   error
}

// Error 实现 error 接口
func (e *ProxyError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s错误: %s (原因: %v)", e.Type, e.Message, e.Cause)
	}
	return fmt.Sprintf("%s错误: %s", e.Type, e.Message)
}

// Unwrap 实现 errors.Unwrap 接口
func (e *ProxyError) Unwrap() error {
	return e.Cause
}

// NewError 创建新的代理错误
func NewError(errType ErrorType, message string, cause error) error {
	return &ProxyError{
		Type:    errType,
		Message: message,
		Cause:   cause,
	}
}

// IsErrorType 检查错误是否为指定类型
func IsErrorType(err error, errType ErrorType) bool {
	if err == nil {
		return false
	}
	if e, ok := err.(*ProxyError); ok {
		return e.Type == errType
	}
	return false
}
